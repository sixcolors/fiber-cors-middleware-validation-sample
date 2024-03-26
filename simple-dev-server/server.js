const express = require('express');

const app = express();
const PORT = process.env.PORT || 8080;

// Serve a simple HTML page
app.get('/', (req, res) => {
    res.send(`
        <html>
            <head>
                <title>Simple Dev Server</title>
            </head>
            <body>
                <h1>Welcome to Simple Dev Server</h1>
                <h2 id="allowedHost">Allowed Host</h2>
                <p>To test Cross-Origin requests, click the buttons below:</p>
                <button id="corsButton">Make CORS GET Request</button>
                <button id="preflightButton">Make CORS POST (Preflight) Request</button>
                <p id="allowedInfo">To test Cross-Origin requests form an <strong>Allowed</strong> host visit <a href="http://localhost:8080/">http://localhost:8080/</a></p>
                <p id="notAllowedInfo">To test Cross-Origin requests form a <strong>Non-Allowed</strong> host visit <a href="http://127.0.0.1:8080/">http://127.0.0.1:8080/</a></p>
                <p>To test Same-Origin GET, POST, and OPTIONS requests visit <a href="http://localhost:3000/">http://localhost:3000/</a></p>
                <script>
                    let allowedHost = document.location.hostname === 'localhost';
                    document.getElementById('corsButton').addEventListener('click', () => {
                        fetch('http://localhost:3000/hello')
                            .then(response => {
                                if (!allowedHost && response.ok) {
									throw new Error('Response is not blocked on Non-Allowed Host for CORS request');
								}
                                return response.text();
                            })
                            .then(data => alert('Response from CORS request: ' + data))
                            .catch(error => {
                                if (error instanceof TypeError && (error.message === 'Failed to fetch' || error.message === 'Load failed')) {
                                    if (!allowedHost) {
                                        alert('Preflight request failed as expected');
                                    } else {
                                        alert('Preflight request failed, but it should not have');
                                    }
                                    return;
                                }
                                alert('Error: ' + error.message);
                            });
                    });
                    document.getElementById('preflightButton').addEventListener('click', () => {
                        fetch('http://localhost:3000/api/1', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({ key: 'value' }),
                        })
                            .then(response => {
                                if (!allowedHost && response.ok) {
									throw new Error('Response is not blocked on Non-Allowed Host for CORS request');
								}
                                return response.json();
                            })
                            .then(data => alert('Response from Preflight request: ' + JSON.stringify(data)))
                            .catch(error => {
                                if (error instanceof TypeError && (error.message === 'Failed to fetch' || error.message === 'Load failed')) {
                                    if (!allowedHost) {
                                        alert('Preflight request failed as expected');
                                    } else {
                                        alert('Preflight request failed, but it should not have');
                                    }
                                    return;
                                }
                                alert('Error: ' + error.message);
                            });
                    });
                    if (!allowedHost) {
                        document.getElementById('allowedHost').innerText = 'Non-Allowed Host';
                        document.getElementById('notAllowedInfo').style.display = 'none';
                    } else {
                        document.getElementById('allowedInfo').style.display = 'none';
                    }
                </script>
            </body>
        </html>
    `);
});

// Start the server
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});

