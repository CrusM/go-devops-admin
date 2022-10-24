package cmd

import "github.com/spf13/cobra"

var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 临时调试代码
// 调试完毕后清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 调试
}
