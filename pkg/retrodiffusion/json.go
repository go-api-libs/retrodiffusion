package retrodiffusion

import (
	"encoding/json/v2"

	"github.com/MarkRosemaker/jsonutil"
)

func init() {
	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true),
		json.WithMarshalers(json.JoinMarshalers(
			json.MarshalToFunc(jsonutil.URLMarshal),
			json.MarshalToFunc(jsonutil.TimeMarshalIntUnix),
		)),
		json.WithUnmarshalers(json.JoinUnmarshalers(
			json.UnmarshalFromFunc(jsonutil.URLUnmarshal),
			json.UnmarshalFromFunc(jsonutil.TimeUnmarshalIntUnix),
		)),
	)
}
