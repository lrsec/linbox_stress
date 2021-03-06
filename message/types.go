package message

type RequestResponseType uint16

const (
	INVALID_REQUEST_RESPONSE_TYPE RequestResponseType = 0
	AUTH_REQUEST_MSG              RequestResponseType = 1
	AUTH_RESPONSE_MSG             RequestResponseType = 2

	SEND_MSG_REQUEST_MSG  RequestResponseType = 9
	SEND_MSG_RESPONSE_MSG RequestResponseType = 10
)

func (rrt RequestResponseType) Name() string {
	switch rrt {
	case INVALID_REQUEST_RESPONSE_TYPE:
		return "INVALID_REQUEST_RESPONSE_TYPE"
	case AUTH_REQUEST_MSG:
		return "AUTH_REQUEST_MSG"
	case AUTH_RESPONSE_MSG:
		return "AUTH_RESPONSE_MSG"
	}

	return ""
}

type MessageType int

const (
	MESSAGE_TYPE_ALL     MessageType = 1
	MESSAGE_TYPE_SESSION MessageType = 2
	MESSAGE_TYPE_GROUP   MessageType = 3
)

// 客户端同步未读信息
//SyncUnreadRequestMsg(3, "SyncUnreadRequest", SyncUnreadRequest.class),
//SyncUnreadResponseMsg(4, "SyncUnreadResponse", SyncUnreadResponse.class),
//
//// 客户端确认未读信息已被读取
//ReadAckRequestMsg(5, "ReadAckRequest", ReadAckRequest.class),
//ReadAckResponseMsg(6, "ReadAckResponse", ReadAckResponse.class),
//
//// 客户端以反序分页拉取信息
//PullOldMsgRequestMsg(7, "PullOldMsgRequest", PullOldMsgRequest.class),
//PullOldMsgResponseMsg(8, "PullOldMsgResponse", PullOldMsgResponse.class),
//

//
//// 新消息通知
//NewMsgInfo(11, "NewMessage", NewMessage.class),
//OfflineInfo(12, "OfflineInfo", OfflineInfo.class),
//
//// Heartbeat
//Ping(13, "Ping", Ping.class),
//Pong(14, "Pong", Pong.class),
//
//// 系统消息
//SystemMsgInfo(15, "SystemMessage", SystemMessage.class),
//SyncSystemUnreadRequestMsg(16, "SyncSystemUnreadRequest", SyncSystemUnreadRequest.class),
//SyncSystemUnreadResponseMsg(17, "SyncSystemUnreadResponse", SyncSystemUnreadResponse.class),
//SystemUnreadAckRequestMsg(18, "SystemUnreadAckRequest", SystemUnreadAckRequest.class),
//SystemUnreadAckResponseMsg(19, "SystemUnreadAckResponse", SystemUnreadAckResponse.class)
