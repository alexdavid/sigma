package mock

import "github.com/alexdavid/sigma"

var mockChats = [][]sigma.Message{
	{
		sigma.Message{FromMe: true, Text: "So a priest, a rabbi, and a chicken walk into a bar."},
		sigma.Message{Text: "I'm pretty sure I've heard this one before"},
		sigma.Message{Text: "So what happens next?"},
		sigma.Message{FromMe: true, Text: "If I remember correctly, they explode outward at the speed of light."},
		sigma.Message{FromMe: true, Text: "But that might be if you cross the streams..."},
		sigma.Message{Text: "... thus negating all existence!"},
		sigma.Message{FromMe: true, Text: "Precisely! it's a risk one takes whenever one walks into a bar, I'm afraid, especially if one is a chicken."},
	}, {
		sigma.Message{Text: "Hi, Peter. What's happening?"},
		sigma.Message{Text: "We need to talk about your TPS reports"},
		sigma.Message{FromMe: true, Text: "Yeah. The coversheet. I know"},
		sigma.Message{FromMe: true, Text: "Bill talked to me about it"},
		sigma.Message{Text: "Yeah. Did you get that memo?"},
		sigma.Message{FromMe: true, Text: "Yeah. I got the memo"},
		sigma.Message{FromMe: true, Text: "And I understand the policy"},
		sigma.Message{FromMe: true, Text: "And the problem is just that I forgot the one time"},
		sigma.Message{FromMe: true, Text: "And I've already taken care of it so it's not even really a problem anymore"},
		sigma.Message{Text: "Ah! Yeah. It's just we're putting new coversheets on all the TPS reports before they go out now"},
		sigma.Message{Text: "So if you could go ahead and try to remember to do that from now on, that'd be great"},
	},
}
