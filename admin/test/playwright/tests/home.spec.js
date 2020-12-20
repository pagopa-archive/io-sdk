describe("test HOME page", () => {
    it("should compare 'Welcome IO-SDK'", async()=> {   
        // open the io-sdk admin
        await page.goto('http://localhost:3280/app/index.html');
        // await page.screenshot({ path: `screenshots/home/home_1.png` })

        const welcome = await page.textContent('text="Welcome"');
        // console.log(welcome);
        expect(welcome).toBe('Welcome');

        const name = await page.textContent('h1');
        // console.log(name);
        expect(name).toBe('IO-SDK');
    })
    it("should compare link to 'GitHub' with right repo link", async()=> {   
        // open the io-sdk admin
        await page.goto('http://localhost:3280/app/index.html');
        // await page.screenshot({ path: `screenshots/home/home_1.png` })

        const linkGitHub = await page.getAttribute('text="GitHub"', 'href');
        expect(linkGitHub).toBe('https://github.com/pagopa/io-sdk');
    })
})