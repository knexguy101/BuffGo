package login

import structs "buffGo/models/client"

//AddLoginDataToClient adds the cookies created on login to the httpclient
func AddLoginDataToClient(l *LoginData, client *structs.HttpClient) {
	for _, v := range l.Cookies {
		client.AddCookie(v.Name, v.Value, v.Path, v.Domain)
	}
}
