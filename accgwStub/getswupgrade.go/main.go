package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sdpon-infra/messaging/client"
	"sdpon-infra/messaging/messages/accessgwmsgs"
	infracommon "sdpon-infra/messaging/messages/common"
)

var Sender client.MsgWriter
var Reader client.MsgReader

func Init() {
	Sender = client.Writer()
	if Sender == nil {
		log.Println("Error in initializing Message Bus of producer")
		os.Exit(1)
	}
	err := Sender.Connect()
	if err != nil {
		log.Println("Error in establishing producer connection to Broker: ", err)
		os.Exit(1)
	}
	Reader = client.Reader()
	Reader.Connect()
}

func msgProcess(msg client.MsgData) {
	if msg != nil {
		switch msg.(type) {
		//switch case with different possibility
		case *client.UserMessage:
			data := msg.(*client.UserMessage)
			writeMsg(data)
		case client.EOFEvent:
			data := msg.(client.EOFEvent)
			log.Println(data.Error)
		case client.ErrorEvent:
			err := msg.(client.ErrorEvent)
			log.Println(err.Str)
		default:
			break
		}
	}
}

func writeMsg(msg *client.UserMessage) {

	msgType := client.GetMessageType(msg)
	// var devMgrMsg devicemgr.AddOltReqv1_0
	switch *msgType {
	case "AddOLTReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.AddOLTReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.AddOLTRespV1_0)
		resp.Status = accessgwmsgs.Success
		resp.SerialNumber = req.SerialNumber
		resp.OLTID = "olt123455678"
		Sender.WriteMsg(resp)
	case "DelOLTReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.DelOLTReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.DelOLTRespV1_0)
		resp.Status = accessgwmsgs.Success
		// resp.SerialNumber = req.SerialNumber
		resp.OLTID = "olt123455678"
		Sender.WriteMsg(resp)
	case "ActivateOLTReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.ActivateOLTReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.ActivateOLTRespV1_0)
		resp.OLTID = "olt123455678"
		resp.Status = accessgwmsgs.Success
		Sender.WriteMsg(resp)
	case "DeactivateOLTReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.DeactivateOLTReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.DeactivateOLTRespV1_0)
		resp.OLTID = "olt123455678"
		resp.Status = accessgwmsgs.Success
		Sender.WriteMsg(resp)

	case "SWDownloadReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.SWDownloadReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.SWDownloadRespV1_0)
		resp.DeviceID = req.DeviceID
		resp.Status = "Accepted"
		Sender.WriteMsg(resp)

		indi := new(accessgwmsgs.SWDownloadIndicationV1_0)
		indi.DeviceID = req.DeviceID
		indi.Status = "Success"
		Sender.WriteMsg(indi)

	case "SWUpdateReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.SWUpdateReqV1_0)
		json.Unmarshal(msg.Val, &req)

		resp := new(accessgwmsgs.SWUpdateRespV1_0)
		resp.DeviceID = req.DeviceID
		resp.Status = "Accepted"
		Sender.WriteMsg(resp)

		// time.Sleep(10000 * time.Millisecond)

		indi := new(accessgwmsgs.SWUpdateIndicationV1_0)
		indi.DeviceID = req.DeviceID
		indi.Status = "Failed"
		Sender.WriteMsg(indi)

		rollback := new(accessgwmsgs.SWRollbackNotifV1_0)
		rollback.DeviceID = req.DeviceID
		rollback.Status = infracommon.SWRollbackInitiated
		Sender.WriteMsg(rollback)

		//rollback = new(accessgwmsgs.SWRollbackNotifV1_0)
		//rollback.DeviceID = req.DeviceID
		//rollback.Status = infracommon.SWRollbackSuccess
		//Sender.WriteMsg(rollback)

		rollback = new(accessgwmsgs.SWRollbackNotifV1_0)
		rollback.DeviceID = req.DeviceID
		rollback.Status = infracommon.SWRollbackFailed
		Sender.WriteMsg(rollback)

	case "SWVersionInfoReq":
		fmt.Println(*msgType)
		req := new(accessgwmsgs.SWVersionInfoReqV1_0)
		json.Unmarshal(msg.Val, &req)
		resp := new(accessgwmsgs.SWVersionInfoRespV1_0)
		resp.DeviceID = req.DeviceID
		resp.PackageInfo = []accessgwmsgs.SWVersionInfo{{ComponentName: "app1", Version: "1.0.0"}}
		Sender.WriteMsg(resp)

	default:
		fmt.Println(*msgType, "not implemneated")
	}
	return
}

func main() {
	Init()
	err := Reader.ReadMsgs([]string{"ACCESSGW"}, msgProcess)
	if err != nil {
		os.Exit(3)
	}
	select {}

}
