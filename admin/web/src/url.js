// parsing url to extract api, menu and key
// expected in format http://<base>#item[/key]
// resolve apihost according different conventions

// url
import {writable} from 'svelte/store'
export const url = writable(location.href)

function parseHash(hash) {
    // check it is not empty
    if(!hash.startsWith("#"))
        return ["/", ""]
    // take the pieces
    let pi  = hash.substring(1).split("/")
    let top = pi.shift()
    return [top, pi.join("/")]
}

// url = new URL("http://user-123-api.noiopen.it/app/index.html")
function parseAPI(url) {
    // handle localhost
    if(url.hostname == "localhost") {
        url.pathname = "/api/v1/web/guest"
        // handle development mode
        if(url.port == 5000)
            url.port = 3280;
        return url.href
    }
    // handle namespace in hostname: if there are 2 hyphens then it is embedded
    let ph = url.hostname.split(".")
    let pn = ph.shift().split("-")
    if(pn.length>2) {
        let ns = pn.shift() + "-" + pn.shift()
        url.hostname = pn.join("-")+"."+ph.join(".")
        url.pathname = "/api/v1/web/"+ns
        return url.href
    }
    // default: use '/api'
    url.pathname = "/api"
    return url.href
}

export function navigate(to, key) {
    let u = new URL(location.href)
    let h = to
    if(key)
      h =+ "#"+key
    u.hash = h
    url.set(u.href)
    location.href = u.href
}

export function parseURL(href) {
    let url = new URL(href)
    // parse hash
    let res = parseHash(url.hash)
    url.hash = ""
    res.push(parseAPI(url))
    return res
}


// samples: 
// parseURL('http://localhost:3280/app/index.html#import/123')
// => [ 'import', '123', 'http://localhost:3280/api/v1/web/guest' ]
// parseURL('http://localhost:3280/app/index.html')
// => [ 'import', '123', 'http://localhost:3280/api/v1/web/guest' ]
// parseURL('http://iosdk.noiopen.it/app/index.html#ship')
// => [ 'ship', '', 'http://iosdk.noiopen.it/api' ]
// parseURL('http://embed-ns-api.noiopen.it/app/index.html#send/123')
// => [ 'send', '123', 'http://api.noiopen.it/api/v1/web/embed-ns' ]

