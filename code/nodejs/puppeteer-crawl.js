const puppeteer = require('puppeteer');
(async () => {
    const browser = await puppeteer.launch({ headless: true });
    const page = await browser.newPage();
    let res1 = await page.goto('https://baidu.cn/gXQQ4RdB', {
        waitUntil: 'networkidle2',
    });
    console.log("Finished");
    const aHandle = await page.evaluateHandle(() => window.shareConfig);
    console.log(await aHandle.jsonValue());

})();