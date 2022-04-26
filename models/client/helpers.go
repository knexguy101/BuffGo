package structs

import (
	"fmt"
	"time"
)

func (c *HttpClient) TS(u string) string {
	return fmt.Sprintf("%s%d", u, time.Now().UnixMilli())
}
