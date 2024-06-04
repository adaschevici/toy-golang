package main

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"log"
	"log/slog"
	"os"
	"product-crawl-extended/internal/eighth_step"
	"product-crawl-extended/internal/fifth_step"
	"product-crawl-extended/internal/first_step"
	"product-crawl-extended/internal/fourth_step"
	"product-crawl-extended/internal/ninth_step"
	"product-crawl-extended/internal/second_step"
	"product-crawl-extended/internal/seventh_step"
	"product-crawl-extended/internal/sixth_step"

	"product-crawl-extended/internal/third_step"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	logger *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.HiBlueString(level)
	case slog.LevelInfo:
		level = color.HiGreenString(level)
	case slog.LevelWarn:
		level = color.HiYellowString(level)
	case slog.LevelError:
		level = color.HiRedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.logger.Println(timeStr, level, msg, color.WhiteString(string(b)))
	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		logger:  log.New(out, "", 0),
	}

	return h
}

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)
	var cmd = &cobra.Command{
		Use:   "crawl",
		Short: "crawl basic set automation1",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("Args being passed in are as follows:", "args", args)
		},
	}
	var basic = &cobra.Command{
		Use:   "basic",
		Short: "crawl basic set automation",
		Long:  "This is the first crawl command and does basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running basic automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			first_step.Crawl()
		},
	}
	var extent_one = &cobra.Command{
		Use:   "extend_one",
		Short: "crawl second stage automation",
		Long:  "This is the second crawl command and does slightly more automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended second automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			second_step.Crawl()
		},
	}
	var extent_two = &cobra.Command{
		Use:   "extend_events",
		Short: "crawl second stage automation with events",
		Long:  "This is the third crawl command and does slightly more automation with events included.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended second automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			third_step.Crawl()
		},
	}
	var extent_three = &cobra.Command{
		Use:   "extend_screenshot",
		Short: "crawl fourth stage automation with taking a screenshot",
		Long:  "This is the fourth crawl command and does slightly more automation with taking a screenshot.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended fourth automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			fourth_step.Crawl()
		},
	}
	var extent_four = &cobra.Command{
		Use:   "extend_form_submission",
		Short: "crawl fifth stage automation with taking a screenshot",
		Long:  "This is the fifth crawl command and does slightly more automation by also sending some key events.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended fifth automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			fifth_step.Crawl()
		},
	}
	var extent_five = &cobra.Command{
		Use:   "extend_static_asset_download",
		Short: "crawl sixth stage automation with taking a screenshot",
		Long:  "This is the sixth crawl command and does slightly more automation by downloading a file that is displayed in the page(ie an image) but does not trigger a background download.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended sixth automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			sixth_step.Crawl()
		},
	}
	var extent_six = &cobra.Command{
		Use:   "extend_file_download",
		Short: "crawl seventh stage automation with taking a screenshot",
		Long:  "This is the seventh crawl command and does slightly more automation by also downloading a file.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended sixth automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			seventh_step.Crawl()
		},
	}
	var extent_seven = &cobra.Command{
		Use:   "extend_set_proxy",
		Short: "crawl eighth stage automation with setting up a proxy and custom User-Agent",
		Long:  "This is the eighth crawl command that sets a proxy and custom User-Agent for the browser.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended sixth automation")
			logger.Debug("Args being passed in are as follows:", "args", args)
			eighth_step.Crawl()
		},
	}

	var extent_eight = &cobra.Command{
		Use:   "extend_blocked",
		Short: "crawl ninth stage but this will be blocked, so we need to figure it out",
		Long:  "This is the ninth crawl command that will get blocked but we should trick it.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("Args being passed in are as follows:", "args", args)
			ninth_step.Crawl()
		},
	}
	rootCmd.AddCommand(cmd)
	cmd.AddCommand(
		basic, extent_one, extent_two, extent_three,
		extent_four, extent_five, extent_six, extent_seven, extent_eight)
	rootCmd.Execute()
}
