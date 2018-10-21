package mock

import (
	"path"
	"runtime"
)

type mockMessage struct {
	FromMe      bool
	Text        string
	Attachments []string
}

func getMockChats() [][]mockMessage {
	return [][]mockMessage{
		{
			{FromMe: true, Text: "So a priest, a rabbi, and a chicken walk into a bar."},
			{Text: "I'm pretty sure I've heard this one before"},
			{Text: "So what happens next?"},
			{FromMe: true, Text: "If I remember correctly, they explode outward at the speed of light."},
			{FromMe: true, Text: "But that might be if you cross the streams..."},
			{Text: "... thus negating all existence!"},
			{FromMe: true, Text: "Precisely! it's a risk one takes whenever one walks into a bar, I'm afraid, especially if one is a chicken."},
		}, {
			{Text: "Hi, Peter. What's happening?"},
			{Text: "We need to talk about your TPS reports"},
			{FromMe: true, Text: "Yeah. The coversheet. I know"},
			{FromMe: true, Text: "Bill talked to me about it"},
			{Text: "Yeah. Did you get that memo?"},
			{FromMe: true, Text: "Yeah. I got the memo"},
			{FromMe: true, Text: "And I understand the policy"},
			{FromMe: true, Text: "And the problem is just that I forgot the one time"},
			{FromMe: true, Text: "And I've already taken care of it so it's not even really a problem anymore"},
			{Text: "Ah! Yeah. It's just we're putting new coversheets on all the TPS reports before they go out now"},
			{Text: "So if you could go ahead and try to remember to do that from now on, that'd be great"},
		}, {
			{Text: "Hey, have you seen my stapler?"},
			{Text: "It looks like this", Attachments: []string{getImgPath("swingline.jpg")}},
		},
	}
}

func getImgPath(imgName string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Mock package must be used in-place")
	}
	return path.Join(path.Dir(fileName), imgName)
}
