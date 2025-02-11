package e

var MsgFlags = map[int]string{
	Success: "ok",
	Error: "fail",
	InvalidParams: "error param",

	ErrorExistUser: "User Name has exited!",
	ErrorFailEncryption: "Failed",
	ErrorExistUserNotFound: "Not Found User",
	ErrorNotCompare:"Password wrong",
	ErrorAuthToken:"Token valid failed",
	ErrorAuthCheckTokenTimeOut: "Token has expired",
	ErrotUploadFail: "failed upload avatar",
	ErrorSendEmail: "Send email failed",
	ErrorProductImgUpload: "Upload picture error",
	ErrorFavoriteExist: "Favorite has existed",
}

//GetMsg code status correspond message
func GetMsg(code int)string{
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}