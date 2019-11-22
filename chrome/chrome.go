package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	U "github.com/neurons-platform/gotools/utils"
	"io/ioutil"
	"log"
	"time"
)

func SavePageAsPng(url string) string {

	var err error

	// // create context
	// ctxt, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // create chrome instance
	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, err := NewHeadless(ctxt, "v5b7.com")
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	//err = c.Run(ctxt, screenshot(`http://v5b7.com/`, `#text-table-of-contents`, &buf))
	// err = c.Run(ctxt, screenshot(url, `#text-table-of-contents`, &buf))
	err = c.Run(ctxt, screenshot(url, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	//err = c.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}

	fileName := U.GetRandomPngFileName()
	err = ioutil.WriteFile("tmp/"+fileName, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}

// NewHeadless 创建headless chrome实例
// chromedp内部有自己的超时设置，你也可以通过ctx来设置更短的超时
func NewHeadless(ctx context.Context, starturl string) (*chromedp.CDP, error) {
	// runner.Flag设置启动headless chrome时的命令行参数
	// runner.URL设置启动时打开的URL
	// Windows用户需要设置runner.Flag("disable-gpu", true)，具体信息参见文档的FAQ
	run, err := runner.New(runner.Flag("headless", true),
		runner.WindowSize(3000, 1500),
		runner.URL(starturl))

	if err != nil {
		return nil, err
	}

	// run.Start启动实例
	err = run.Start(ctx)
	if err != nil {
		return nil, err
	}

	// 默认情况chromedp会输出大量log，因为是示例所以选择屏蔽
	// 使用runner初始化chromedp实例
	// 实例在使用完毕后需要调用c.Shutdown()来释放资源
	c, err := chromedp.New(ctx, chromedp.WithRunner(run), chromedp.WithErrorf(log.Printf))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func screenshot(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(10 * time.Second),
		chromedp.CaptureScreenshot(res),
	}
}
