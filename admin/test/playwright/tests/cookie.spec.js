describe("test", () => {
    it("accepts cookies", async()=> {
        await context.clearCookies()
        let cookies = await context.cookies()
        //console.log(cookies)
   
        await page.goto('http://localhost:3280/app/index.html');
        let accept = await page.innerText(".cookiebar button");
        expect(accept).toBe("Accetto")
       
        await page.click('text="Accetto"'); 
        cookies = await context.cookies()
        //console.log(cookies)
        expect(cookies[0]['name']).toBe('cookies_consent')
    })
})