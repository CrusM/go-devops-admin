package cmd

import (
	"go-devops-admin/pkg/cache"
	"go-devops-admin/pkg/console"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "cache management",
}

var cacheKey string

func init() {
	CmdCache.AddCommand(
		CmdCacheClear,
		CmdCacheForget,
	)

	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY for the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "clear all cache key, cache.Flush()",
	Run:   runCacheClear,
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("cache cleared!")
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete cache key, example: cache forget cache-key",
	Run:   runCacheForget,
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
}
