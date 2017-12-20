package pili

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Meeting struct {
	rpc *RPC
}

func NewMeeting(creds *Credentials) Meeting {
	return Meeting{rpc: NewRPC(creds)}
}

func (c Meeting) CreateRoom(args RoomOptionArguments) (room Room, err error) {
	if args.OwnerId == "" {
		err = errors.New("illegal owner")
		return
	}
	data := map[string]interface{}{"owner_id": args.OwnerId}
	if args.Name != "" {
		data["room_name"] = args.Name
	}
	if args.UserMax > 0 {
		data["user_max"] = args.UserMax
	}
	url := fmt.Sprintf("%s/rooms", getRtcApiBaseUrl())
	err = c.rpc.PostCall(&room, url, data)
	if err != nil {
		return
	}
	return room, nil
}

func (c Meeting) GetRoom(roomName string) (room Room, err error) {
	if roomName == "" {
		err = errors.New("illegal room")
		return
	}
	url := fmt.Sprintf("%s/rooms/%s", getRtcApiBaseUrl(), roomName)
	err = c.rpc.GetCall(&room, url)
	return
}

func (c Meeting) ActiveUsers(roomName string) (allActiveUsers AllActiveUsers, err error) {
	if roomName == "" {
		err = errors.New("illegal room")
		return
	}
	url := fmt.Sprintf("%s/rooms/%s/users", getRtcApiBaseUrl(), roomName)
	err = c.rpc.GetCall(&allActiveUsers, url)
	return
}

func (c Meeting) RejectUser(roomName, userId string) (err error) {
	if roomName == "" {
		err = errors.New("illegal room")
		return
	}
	if userId == "" {
		err = errors.New("illegal parameter")
		return
	}
	url := fmt.Sprintf("%s/rooms/%s/users/%s", getRtcApiBaseUrl(), roomName, userId)
	var ret interface{}
	err = c.rpc.DelCall(&ret, url)
	return
}

func (c Meeting) RoomToken(creds *Credentials, args RoomAccessPolicy) (token string, err error) {
	data, pErr := json.Marshal(args)
	if pErr != nil {
		err = pErr
		return
	}
	token = creds.SignWithData(data)
	return
}

func (c Meeting) DeleteRoom(roomName string) (ret interface{}, err error) {
	if roomName == "" {
		err = errors.New("illegal room")
		return
	}
	url := fmt.Sprintf("%s/rooms/%s", getRtcApiBaseUrl(), roomName)
	err = c.rpc.DelCall(&ret, url)
	return
}
