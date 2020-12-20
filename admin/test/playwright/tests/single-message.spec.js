describe("test Single Message function", () => {
    // to fix timeout
    xit("should compare 'loading...'", async()=> {
        await page.route('**/api/v1/web/guest/util/cache', route => route.fulfill({
            status: 200,
            body: '{"scan": ["message:AAAAAA00A00A000A:1"]}',
          }));

        // open the io-sdk admin
        await page.goto('http://localhost:3280/app/index.html');
        // await page.screenshot({ path: `screenshots/single-message/loading_1.png` })

        // navigate to page single message
        await page.click('a[href="#send"]');
        // await page.screenshot({ path: `screenshots/single-message/loading_2.png` })

        const loading = await page.innerText('text="loading..."');
        console.log(loading);
        expect(loading).not.toBe(null);
    })

    describe("test FISCAL CODE field", () => {
        it("I enter empty fiscal code, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', '');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#fiscal_code', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter wrong fiscal code format, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'AAAA');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#fiscal_code', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right fiscal code, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#fiscal_code', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })

    describe("test SUBJECT field", () => {
        it("I enter empty in subject, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', '');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#subject', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter more than 255 chars in subject, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'A'.repeat(256));
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#subject', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter less than 5 chars in subject, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#subject', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right subject, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#subject', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })

    describe("test MARKDOWN field", () => {
        it("I enter empty markdown, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', '');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#markdown', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right markdown, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#markdown', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })

    describe("test AMOUNT field", () => {
        it("I enter empty amount, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '');
            await page.fill('#notice_number', '123');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#amount', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter not numeric amount, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', 'AAA');
            await page.fill('#notice_number', '123');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#amount', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right amount, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#amount', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })

    describe("test NOTICE NUMBER field", () => {
        it("I enter empty notice_number, other field ok, and should set it as invalid and button is disabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');        
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '');
            // // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#notice_number', 'class');
            expect(inputSUTClass).toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe("");
        })
        it("I enter right notice number, other field ok, and should set it as valid and button is enabled", async()=> {   
            // open the io-sdk admin
            await page.goto('http://localhost:3280/app/index.html');
            // await page.screenshot({ path: `screenshots/single-message/send_1.png` })

            // navigate to page single message
            await page.click('a[href="#send"]');
            // // await page.screenshot({ path: `screenshots/single-message/send_2.png` })

            await page.fill('#fiscal_code', 'ABCDEF12G34J456T');
            await page.fill('#subject', 'AAAAA');
            await page.fill('#markdown', 'AAAAA');
            await page.fill('#amount', '123');
            await page.fill('#notice_number', '123');
            // await page.screenshot({ path: `screenshots/single-message/send_3.png` })

            const inputSUTClass = await page.getAttribute('#notice_number', 'class');
            expect(inputSUTClass).not.toMatch(/is-invalid/);

            const buttonDisabled = await page.getAttribute('button', 'disabled');
            expect(buttonDisabled).toBe(null);
        })
    })
})