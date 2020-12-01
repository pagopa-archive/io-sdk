const GITHUB_HOST = 'https://api.github.com';
const GITHUB_RELEASES_ENDPONT = '/repos/pagopa/io-gateway-connectors/releases';
const GITHUB_LATEST_RELEASE_ENDPOINT = '/repos/pagopa/io-gateway-connectors/releases/latest';

const getIoGetawayConnectorReleases = () => {

    const url = GITHUB_HOST + GITHUB_RELEASES_ENDPONT;

    return new Promise((reject,resolve) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => {

            const releases = await res.json();
            resolve( releases );

        })
        .catch( err => reject(err) )

    })

}

const getIoGetawayConnectorLatestRelease = () => {

    const url = GITHUB_HOST + GITHUB_LATEST_RELEASE_ENDPOINT;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => {

            const latestRelease = await res.json();
            resolve( latestRelease );

        })
        .catch( err => reject(err) )

    })

}

export const githubService = { getIoGetawayConnectorReleases, getIoGetawayConnectorLatestRelease }