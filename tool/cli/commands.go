// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": CLI Commands
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package cli

import (
	"EvelyApi/client"
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// SigninAuthCommand is the command line data structure for the signin action of auth
	SigninAuthCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SignupAuthCommand is the command line data structure for the signup action of auth
	SignupAuthCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// CreateEventsCommand is the command line data structure for the create action of events
	CreateEventsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteEventsCommand is the command line data structure for the delete action of events
	DeleteEventsCommand struct {
		// イベントID
		EventID string
		// ユーザーID
		UserID      string
		PrettyPrint bool
	}

	// ListEventsCommand is the command line data structure for the list action of events
	ListEventsCommand struct {
		// キーワード
		Keyword string
		// 取得件数
		Limit int
		// 除外件数
		Offset      int
		PrettyPrint bool
	}

	// ShowEventsCommand is the command line data structure for the show action of events
	ShowEventsCommand struct {
		// イベントID
		EventID string
		// ユーザーID
		UserID      string
		PrettyPrint bool
	}

	// UpdateEventsCommand is the command line data structure for the update action of events
	UpdateEventsCommand struct {
		Payload     string
		ContentType string
		// イベントID
		EventID string
		// ユーザーID
		UserID      string
		PrettyPrint bool
	}

	// ShowUsersCommand is the command line data structure for the show action of users
	ShowUsersCommand struct {
		// ユーザーID
		UserID      string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `イベント作成`,
	}
	tmp1 := new(CreateEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/develop/v1/events"]`,
		Short: ``,
		Long: `

Payload example:

{
   "body": "初心者でもGitを扱えるようになる勉強会を開催します！\nノートPCを各自持参してください。",
   "mail": "yKicchanApp@gmail.com",
   "place": {
      "lat": 34.706424,
      "lng": 135.50123,
      "name": "ECCコンピュータ専門学校2303教室"
   },
   "tel": "090-1234-5678",
   "title": "Git勉強会",
   "upcomingDate": {
      "endDate": "1974-02-03T21:26:56Z",
      "startDate": "2003-06-21T02:17:51Z"
   },
   "url": "http://comp.ecc.ac.jp/"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `イベント削除`,
	}
	tmp2 := new(DeleteEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/develop/v1/events/USER_ID/EVENT_ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `イベント複数取得`,
	}
	tmp3 := new(ListEventsCommand)
	sub = &cobra.Command{
		Use:   `events [("/api/develop/v1/events"|"/api/develop/v1/events/USER_ID")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp4 := new(ShowEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/develop/v1/events/USER_ID/EVENT_ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(ShowUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/develop/v1/users/USER_ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `ログイン`,
	}
	tmp6 := new(SigninAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/develop/v1/auth/signin"]`,
		Short: ``,
		Long: `

Payload example:

{
   "id": "yKicchan",
   "password": "password"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signup",
		Short: `新規登録`,
	}
	tmp7 := new(SignupAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/develop/v1/auth/signup"]`,
		Short: ``,
		Long: `

Payload example:

{
   "id": "yKicchan",
   "mail": "yKicchanApp@gmail.com",
   "name": "きっちゃそ",
   "password": "password",
   "tel": "090-1234-5678"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `イベント編集`,
	}
	tmp8 := new(UpdateEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/develop/v1/events/USER_ID/EVENT_ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "body": "初心者でもGitを扱えるようになる勉強会を開催します！\nノートPCを各自持参してください。",
   "mail": "yKicchanApp@gmail.com",
   "place": {
      "lat": 34.706424,
      "lng": 135.50123,
      "name": "ECCコンピュータ専門学校2303教室"
   },
   "tel": "090-1234-5678",
   "title": "Git勉強会",
   "upcomingDate": {
      "endDate": "1974-02-03T21:26:56Z",
      "startDate": "2003-06-21T02:17:51Z"
   },
   "url": "http://comp.ecc.ac.jp/"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swaggerui/") {
		fnd = c.DownloadSwaggerui
		rpath = rpath[11:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the SigninAuthCommand command.
func (cmd *SigninAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/develop/v1/auth/signin"
	}
	var payload client.LoginPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninAuth(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SignupAuthCommand command.
func (cmd *SignupAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/develop/v1/auth/signup"
	}
	var payload client.UserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SignupAuth(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignupAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the CreateEventsCommand command.
func (cmd *CreateEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/develop/v1/events"
	}
	var payload client.EventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateEvents(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteEventsCommand command.
func (cmd *DeleteEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/develop/v1/events/%v/%v", url.QueryEscape(cmd.UserID), url.QueryEscape(cmd.EventID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteEvents(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID string
	cc.Flags().StringVar(&cmd.EventID, "event_id", eventID, `イベントID`)
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "user_id", userID, `ユーザーID`)
}

// Run makes the HTTP request corresponding to the ListEventsCommand command.
func (cmd *ListEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/develop/v1/events"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListEvents(ctx, path, cmd.Limit, cmd.Offset, stringFlagVal("keyword", cmd.Keyword))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var keyword string
	cc.Flags().StringVar(&cmd.Keyword, "keyword", keyword, `キーワード`)
	cc.Flags().IntVar(&cmd.Limit, "limit", 10, `取得件数`)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, `除外件数`)
}

// Run makes the HTTP request corresponding to the ShowEventsCommand command.
func (cmd *ShowEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/develop/v1/events/%v/%v", url.QueryEscape(cmd.UserID), url.QueryEscape(cmd.EventID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowEvents(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID string
	cc.Flags().StringVar(&cmd.EventID, "event_id", eventID, `イベントID`)
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "user_id", userID, `ユーザーID`)
}

// Run makes the HTTP request corresponding to the UpdateEventsCommand command.
func (cmd *UpdateEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/develop/v1/events/%v/%v", url.QueryEscape(cmd.UserID), url.QueryEscape(cmd.EventID))
	}
	var payload client.EventPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateEvents(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var eventID string
	cc.Flags().StringVar(&cmd.EventID, "event_id", eventID, `イベントID`)
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "user_id", userID, `ユーザーID`)
}

// Run makes the HTTP request corresponding to the ShowUsersCommand command.
func (cmd *ShowUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/develop/v1/users/%v", url.QueryEscape(cmd.UserID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUsers(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "user_id", userID, `ユーザーID`)
}
