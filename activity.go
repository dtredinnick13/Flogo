/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package tcmpublisher

import (
	"fmt"

	"github.com/TIBCOSoftware/eftl"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/project-flogo/core/engine"

	"encoding/json"
	"sync"

	"git.tibco.com/git/product/ipaas/wi-contrib.git/connection/generic"
)

const (
	ivConnection = "tcmConnection"
	ivMessage    = "message"
	//ivUrl        = "url"
	//ivAuthKey    = "authKey"
	ivDest = "destination"
)

var activityLog = logger.GetLogger("tcm-activity-messagepublisher")

var cache *ConnectionCache

type TCMSendMessageActivity struct {
	metadata *activity.Metadata
}

type ConnectionCache struct {
	connMap map[string]*eftl.Connection
	lock    *sync.RWMutex
}

type CacheManager struct {
}

func (managedCache CacheManager) Start() error {
	return nil
}

func (managedCache CacheManager) Stop() error {

	if cache != nil {
		for _, conn := range cache.connMap {
			if conn.IsConnected() {
				conn.Disconnect()
			}
		}
	}
	return nil
}

func (cache *ConnectionCache) AddConnection(id string, conn *eftl.Connection) {
	defer cache.lock.Unlock()
	cache.lock.Lock()
	cache.connMap[id] = conn
}

func (cache *ConnectionCache) GetConnection(id string) *eftl.Connection {
	defer cache.lock.RUnlock()
	cache.lock.RLock()
	return cache.connMap[id]
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	engine.LifeCycle(CacheManager{})
	cache = &ConnectionCache{lock: &sync.RWMutex{}, connMap: make(map[string]*eftl.Connection)}
	return &TCMSendMessageActivity{metadata: metadata}
}

func (a *TCMSendMessageActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *TCMSendMessageActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Debug("Executing Message Publisher activity")

	if context.GetInput(ivMessage) == nil {
		return false, activity.NewError("Message must be configured", "TCM-SENDMESSAGE-4002", nil)
	}

	message := make(map[string]interface{})

	//Read mapped values
	messageAttributes, _ := data.CoerceToComplexObject(context.GetInput(ivMessage))
	if messageAttributes != nil && messageAttributes.Value != nil {
		switch messageAttributes.Value.(type) {
		case map[string]interface{}:
			message, _ = data.CoerceToObject(messageAttributes.Value)
		}
	}

	destName, _ := data.CoerceToString(context.GetInput(ivDest))
	if len(destName) > 0 {
		message["_dest"] = destName
	}

	jsonString, parseErr := json.Marshal(message)

	if parseErr != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to convert the message body to string due to error - {%s}.", parseErr.Error()), "TCM-MESSAGEPUB-4004", nil)
	}

	errChan := make(chan error, 1)

	id := context.ActivityHost().Name() + context.Name()

	conn := cache.GetConnection(id)

	if conn == nil {

		//Read Inputs
		connection, err := generic.NewConnection(context.GetInput(ivConnection))
		if err != nil {
			return false, activity.NewError("Connection must be configured", "TCM-MESSAGEPUB-4001", nil)
		}
		url := connection.GetSetting("url").(string)
		authkey := connection.GetSetting("authKey").(string)

		// set connection options
		opts := &eftl.Options{
			Password: authkey,
		}

		activityLog.Debug("Connecting to TIBCO Cloud Messaging Service")

		// connect to TIBCO Cloud Messaging
		conn, err = eftl.Connect(url, opts, errChan)
		if err != nil {
			return false, activity.NewError(fmt.Sprintf("Failed to connect to TIBCO Cloud Messaging service due to error - {%s}. Either connection parameters are invalid or messaging service is down.", err.Error()), "TCM-MESSAGEPUB-4004", nil)
		}
		cache.AddConnection(id, conn)
	}

	// publish a message to TIBCO Cloud Messaging
	eftlMsg := eftl.Message{}
	err = eftlMsg.UnmarshalJSON([]byte(jsonString))
	if err != nil {
		return false, err
	}
	err = conn.Publish(eftlMsg)
	activityLog.Debug(eftlMsg)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to send message to TIBCO Cloud Messaging service due to error - {%s}.", err.Error()), "TCM-MESSAGEPUB-4005", nil)
	}
	activityLog.Info("Message published")
	return true, nil
}
