const GITHUB_REPO_HOST  = 'https://sf3ris.github.io';
const GITHUB_REPO_NAME  = '/io-gateway-connectors/';
const GITHUB_REPO_INDEX = 'index.json';

const getIoGetawayConnectors = () => {

    const url = GITHUB_REPO_HOST + GITHUB_REPO_NAME + GITHUB_REPO_INDEX;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => {

            resolve(await res.json());

        })
        .catch( err => reject(err) )

    })

}

const getIoGetawayConnectorBinary = ( binaryName ) => {
    
    const url = GITHUB_REPO_HOST + GITHUB_REPO_NAME + binaryName;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => {

            resolve(await res.json());

        })
        .catch( err => reject(err) )

    })

}

export const githubService = { getIoGetawayConnectors, getIoGetawayConnectorBinary }