package pili

import (
	"encoding/json"
	"fmt"
)

type Meeting struct {
	rpc     *RPC
	ownerId string
}

func NewMeeting(creds *Credentials, ownerId string) Meeting {
	return Meeting{rpc: NewRPC(creds), ownerId: ownerId}
}

func (c Meeting) CreateRoom(args RoomOptionArguments) (room Room, err error) {
	data := map[string]interface{}{"owner_id": c.ownerId}
	if args.Name != "" {
		data["room_name"] = args.Name
	}
	if args.UserMax > 0 {
		data["user_max"] = args.UserMax
	}
	if args.Version != "" {
		data["version"] = args.Version
	}
	url := fmt.Sprintf("%s/rooms", getRtcApiBaseUrl())
	err = c.rpc.PostCall(&room, url, data)
	if err != nil {
		return
	}
	return room, nil
}

func (c Meeting) GetRoom(roomName string) (room Room, err error) {
	url := fmt.Sprintf("%s/rooms/%s", getRtcApiBaseUrl(), roomName)
	err = c.rpc.GetCall(&room, url)
	if err != nil {
		return
	}
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
	url := fmt.Sprintf("%s/rooms/%s", getRtcApiBaseUrl(), roomName)
	err = c.rpc.DelCall(&ret, url)
	return
}
