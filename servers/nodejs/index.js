const express = require('express');
const auth = require('basic-auth')
const app = express();
app.use(express.json());

const PORT = 3060;

function main() {
    const username = process.env.USERNAME || 'username';
    const password = process.env.PASSWORD || 'password';


    app.post('/', (request, response) => {
        const workspaceId = request.header('Plain-Workspace-Id');
        const plainEventType = request.header('Plain-Event-Type');

        if (!workspaceId) {
            console.log('No Plain-Workspace-Id header found')
            response.status(400)
            response.json({message: 'Bad request'})
            return;
        }

        if (!plainEventType) {
            console.log('No Plain-Event-Type header found')
            response.status(400)
            response.json({message: 'Bad request'})
            return;
        }

        const user = auth(request)
        if (!user) {
            console.log('No auth credentials were provided in the request')
            response.status(401)
            response.json({message: 'Unauthorized'})
            return;
        }

        if (user.name !== username || user.pass !== password) {
            console.log('Either the username or password do not match')
            response.status(401)
            response.json({message: 'Unauthorized'})
            return
        }

        console.log(`Received ${plainEventType} event from workspace ${workspaceId}: %o`, request.body);
        response.json({message: 'ok'});
    });

    app.listen(PORT, () => {
        console.log(`Webhook handler running on http://localhost:${PORT}`);
    });
}

main()