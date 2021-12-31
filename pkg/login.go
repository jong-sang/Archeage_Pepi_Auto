package pkg

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func Login(AccType string, LoginID string, LoginPW string) error {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.DisableGPU,
		// chromedp.UserDataDir("someUserDir"),
		chromedp.Flag("headless", false),
		// chromedp.Flag("enable-automation", false),
		// chromedp.Flag("no-default-browser-check", true),
		// chromedp.Flag("restore-on-startup", false),
		// chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"),
	)

	contextVar, cancelFunc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelFunc()

	contextVar, cancelFunc = chromedp.NewContext(contextVar)
	defer cancelFunc()

	ch := chromedp.WaitNewTarget(contextVar, func(info *target.Info) bool {
		return info.URL != ""
	})
	if err := chromedp.Run(contextVar,
		chromedp.Navigate(`https://www.naver.com/`), // 시작 URL
		chromedp.SendKeys(`query`, `아키에이지`, chromedp.ByID),
		chromedp.Sleep(time.Second*1),
		chromedp.Click(`search_btn`, chromedp.ByID),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`document.querySelector("#main_pack > section.sc_new.cs_common_module._cs_newgame.case_normal.color_10 > div.cm_content_wrap.cm_game > div:nth-child(1) > div > div.detail_info > dl > div:nth-child(5) > dd > a:nth-child(2)")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Second*2),
		// chromedp.InnerHTML(`html`, &html),

		// chromedp.Click(`document.querySelector("#account_util > div.xlgames-login > a")`, chromedp.ByJSPath),
		// chromedp.Sleep(time.Second*2),
	); err != nil {
		return err
	}

	newCtx, cancel := chromedp.NewContext(contextVar, chromedp.WithTargetID(<-ch))
	defer cancel()
	var urlstr string
	// var x string

	if err := chromedp.Run(newCtx,
		chromedp.Location(&urlstr),
		chromedp.Click(`document.querySelector("#account_util > div.xlgames-login > a")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Second*2),
		chromedp.SendKeys(`id_field`, LoginID, chromedp.ByID),
		chromedp.SendKeys(`pw_field`, LoginPW, chromedp.ByID),
		chromedp.Click(`loginButton`, chromedp.ByID),
		chromedp.Sleep(time.Second*1),
		chromedp.Navigate(`https://archeage.xlgames.com/events/running/679`),
		chromedp.Sleep(time.Second*1),
		chromedp.Click(`.link-gift`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*5),
	); err != nil {
		return err
	}

	return nil
}
