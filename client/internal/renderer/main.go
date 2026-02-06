package renderer

import (
	"fmt"

	"github.com/persona-mp3/client/common"
)

// From a message-type of ContentPaintType from server
// we get to display all the friends the user has
// Similar to how you open instagram or whatsapp
func RenderContentPaint(friends *[]common.Friend) error {
	for i, user := range *friends {
		fmt.Printf("   %1d. %20s | ", i+1, user.Name)
		fmt.Printf("lastActive: %2s\n", user.LastSeen)
	}

	return nil
}

func RenderChatMessage(msg *common.Message) {
	fmt.Printf("  [%s]:  %s \n", msg.From, msg.Message)
}
