const IOSDK_HOST = 'http://localhost:3280/api/v1';
const IOSDK_NAMESPACE = 'guest';

const createConnectors = ( connectors ) => {

    const url = IOSDK_HOST + `/web/${IOSDK_NAMESPACE}/util/select_connectors?blocking=true&result=true`;

    const data = {
        connectors
    }

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'PUT',
                body: JSON.stringify(data),
                headers: {
                    "Content-Type": "application/json"
                }
            }
        )
        .then( async res => resolve(await res.json()))
        .catch( err => reject(err))

    })

}

const getCustomConnectors = () => {

    const url = IOSDK_HOST + `/web/${IOSDK_NAMESPACE}/util/get_custom_connectors?blocking=true&result=true`;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => resolve(await res.json()))
        .catch( err => reject(err))

    })

}

const getCustomConnectorParams = ( connector_name ) => {

    const url = IOSDK_HOST + `/web/${IOSDK_NAMESPACE}/util/get_custom_connector_params?blocking=true&result=true&connector_name=${connector_name}`;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'GET'
            }
        )
        .then( async res => resolve(await res.json()))
        .catch( err => reject(err))

    })

}

const triggerCustomConnectorAction = (api, connector_name, params ) => {

    const url = `${api}/util/${connector_name}`;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'POST',
                body: JSON.stringify(params),
                headers: {
                    "Content-Type": "application/json"
                }
            }
        )
        .then( async res => resolve(await res.json()))
        .catch( err => reject(err))
    })

}

const storeMessages = (api, messages ) => {

    const url = `${api}/util/store`;

    return new Promise((resolve,reject) => {

        fetch(
            url,
            {
                method: 'POST',
                body: JSON.stringify({data: messages}),
                headers: {
                    "Content-Type": "application/json"
                }
            }
        )
        .then( async res => resolve(await res.json()))
        .catch( err => reject(err))

    });

}

export const ioSDKService = { 
    createConnectors, 
    getCustomConnectors, 
    getCustomConnectorParams, 
    triggerCustomConnectorAction,
    storeMessages
}