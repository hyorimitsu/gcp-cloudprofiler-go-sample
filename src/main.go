package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/profiler"
)

func busyloop() {
	for i := 0; i < 1000; i++ {
		foo(1000)
	}
}

func foo(scale int) {
	bar(scale)
	baz(scale)
}

func bar(scale int) {
	load(scale)
}

func baz(scale int) {
	load(scale)
}

func load(scale int) {
	for i := 0; i < scale*(1<<14); i++ {
	}
}

func main() {

	// GAE上で動いている場合のみプロファイリング設定（GAE上で動かした場合は GOOGLE_CLOUD_PROJECT に ProjectID が設定されている）
	if os.Getenv("GOOGLE_CLOUD_PROJECT") != "" {
		err := profiler.Start(profiler.Config{
			// サービス名。デフォルトはGAEに存在する GAE_SERVICE（app.yaml の service: service_name。それも無い場合は default）。
			Service:              os.Getenv("SERVICE_NAME"),
			// バージョン。デフォルトはGAEに存在する GAE_VERSION。過去のバージョンとプロファイルを比較できる。
			ServiceVersion:       os.Getenv("SERVICE_VERSION"),
			// プロファイラからのデバッグログを出力する場合に指定。デフォルトはfalse。
			//DebugLogging:         true,
			// 排他ロックによる待機時間をプロファイルする場合に指定。デフォルトはfalse。
			//MutexProfiling:       false,
			// CPU時間のプロファイルを無効にする場合に指定。デフォルトはfalse。
			//NoCPUProfiling:       false,
			// 割り当てられたヒープのプロファイルを無効にする場合に指定。デフォルトはfalse。
			//NoAllocProfiling:     false,
			// ヒープのプロファイルを収集する前にGCを強制し、「Allocated heap」を生成する場合に指定。
			// これにより「Allocated heap」の精度が向上する。デフォルトはfalse。
			//AllocForceGC:         false,
			// ヒープのプロファイルを無効にする場合に指定。デフォルトはfalse。
			//NoHeapProfiling:      false,
			// goroutineのプロファイルを無効にする場合に指定。デフォルトはfalse。
			//NoGoroutineProfiling: false,
			// OpenCensusエクスポーターを介してすべてのテレメトリを送信する場合に指定。デフォルトはfalse。
			//EnableOCTelemetry:    false,
			// ローカル環境で実行する場合にGCPのProjectIDを指定。デフォルトはGAEに存在する GOOGLE_CLOUD_PROJECT。
			//ProjectID:            "",
			// プロファイラエージェントAPIへの接続に使用するHTTPエンドポイント。デフォルトは「cloudprofiler.googleapis.com:443」。
			// テストのためにオーバライドが可能。
			//APIAddr:              "",
			// プロファイラエージェントが実行されているGCEインスタンスの名前。
			// 通常はGCEメタデータサーバーから決定されるため設定不要。
			// メタデータサーバーが不安定な場合などの稀なケースで利用。
			//Instance:             "",
			// プロファイラエージェントが実行されているGCEインスタンスのゾーン。
			// 通常はGCEメタデータサーバーから決定されるため設定不要。
			// メタデータサーバーが不安定な場合などの稀なケースで利用。
			//Zone:                 "",
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to start the profiler: %v", err))
		}
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	http.HandleFunc("/heavy", func (w http.ResponseWriter, r *http.Request) {
		busyloop()
		fmt.Fprintf(w, "SUCCESS: heavy")
	})
	http.HandleFunc("/sleep", func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(600 * time.Second)
		fmt.Fprintf(w, "SUCCESS: sleep")
	})

	// GAE上で動かした場合は PORT が設定されている
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)

}