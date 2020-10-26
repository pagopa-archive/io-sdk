const { firefox } = require('playwright');
(async () => {
  const browser = await firefox.launch({
    slowMo: 100,
    headless: false
  });
  const context = await browser.newContext();

  // Open new page
  const page = await context.newPage();

  // Go to http://localhost:3280/app/index.html
  await page.goto('http://localhost:3280/app/index.html');

  // Click text="Accetto"
  await page.click('text="Accetto"');

  // Click //a[normalize-space(.)='Import URL']
  await page.click('//a[normalize-space(.)=\'Import URL\']');
  // assert.equal(page.url(), 'http://localhost:3280/app/index.html#import');

  // Fill input[name="username"]
  await page.fill('input[name="username"]', 'postgres');

  // Fill input[name="password"]
  await page.fill('input[name="password"]', 'postgrespassword');
  
  await page.fill('textarea[name="jsonargs"]', '{"count":20,"fiscal_code":"AAAAAA00A00A000A","amount":0,"due_date":"","notice_number":""}');

  await page.click('//button[normalize-space(.)=\'Import\']', 'Enter');

  // Click //a[normalize-space(.)='Send Messages']
  await page.click('//a[normalize-space(.)=\'Send Messages\']');

  // Click text="Select All"
  await page.click('text="Select All"');

  const handle = await page.$('text="OK"');
  console.log(handle)

  // Click text="Send Selected Messages"
  await page.click('text="Send Selected Messages"');

  // ---------------------
  await context.close();
  await browser.close();
})();
