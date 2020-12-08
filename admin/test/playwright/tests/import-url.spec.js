const jsonArgs = '{"count":1,"fiscal_code":"AAAAAA00A00A000A","amount":0,  "due_date":"", "notice_number":""}';

describe("test Import URL function", () => {
    // fix it - timeout, maybe retrieve data is too fast and isLoading is gone
    xit("should compare 'loading...'", async()=> {   
        // open the io-sdk admin
        await page.goto('http://localhost:3280/app/index.html');
        await page.screenshot({ path: `screenshots/importURL/loading_1.png` })

        // navigate to page single message
        await page.click('a[href="#import"]');
        let loading = await page.innerText('text="loading..."');
        await page.screenshot({ path: `screenshots/importURL/loading_2.png` })

        console.log(loading);
        expect(loading).toBe('loading...');
    })
    // timeout when I click button
    xit("I enter data and submit form", async()=> {   
        // open the io-sdk admin
        await page.goto('http://localhost:3280/app/index.html');
        await page.screenshot({ path: `screenshots/importURL/loading_1.png` })

        // navigate to page single message
        await page.click('a[href="#import"]');
        await page.screenshot({ path: `screenshots/importURL/loading_2.png` })

        await page.fill('#url', 'AAAA');
        await page.fill('#username', 'AAAA');
        await page.fill('#password', 'AAAA');
        await page.fill('#jsonargs', jsonArgs);
        await page.uncheck('#useget');
        await page.click('button');
        // let loading = await page.innerText('text="loading..."');
        // await page.screenshot({ path: `screenshots/importURL/loading_3.png` })

        // console.log(loading);
        // expect(loading).not.toBe(null);

        // let tableImport = await page.awaitForSelector('table');
        // console.table(tableImport);
        // expect(tableImport).not.toBe(null);


        // const invalidClass = await page.$('#fiscalcode inInvalid');
        // console.log(invalidClass)
        // expect(invalidClass).toBe('loading...');
    })

    describe("test URL field", () => {
        it("I enter empty url, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/import-url/form_1.png` })

            // navigate to page import url
            await page.click('a[href="#import"]');
            // await page.screenshot({ path: `screenshots/import-url/form_2.png` })

            await page.fill('#url', '');
            await page.fill('#username', 'demo');
            await page.fill('#password', 'demo');
            await page.fill('#jsonargs', jsonArgs);
            await page.uncheck('#useget');
            // // await page.screenshot({ path: `screenshots/import-url/form_3.png` })

            const inputSUTClass = await page.getAttribute('#url', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right URL, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/import-url/form_1.png` })

            // navigate to page single message
            await page.click('a[href="#import"]');
            // // await page.screenshot({ path: `screenshots/import-url/form_2.png` })

            await page.fill('#url', 'http://localhost:3280/api/v1/web/guest/util/sample');
            await page.fill('#username', 'AAAA');
            await page.fill('#password', 'AAAA');
            await page.fill('#jsonargs', jsonArgs);
            await page.uncheck('#useget');
            // await page.screenshot({ path: `screenshots/import-url/form_3.png` })

            const inputSUTClass = await page.getAttribute('#url', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })
})