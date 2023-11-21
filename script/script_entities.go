package script

type OptScr func(*ScriptOpts)

type Script struct {
	ScriptOpts
}

type ScriptOpts struct {
	CreatedAt string `json:"created_at"`
	Event     string `json:"event"`
	ID        int    `json:"id"`
	Src       string `json:"src"`
	UpdatedAt string `json:"updated_at"`
	Where     string `json:"where"`
}
