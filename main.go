package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:8080",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.SetContentType("text/html")
		return c.SendString(`
        <html>
            <head>
                <title>CORS Test Server</title>
            </head>
            <body>
                <h1>Welcome to CORS Test Server</h1>
                <p>You can test same-origin GET, POST, and OPTIONS requests here.</p>
                <button id="getButton">Make GET Request</button>
                <button id="postButton">Make Post Request</button>
				<button id="optionsButton">Make OPTIONS Request (Non-CORS)</button>
                <p>To test Cross-Origin requests form an <strong>Allowed</strong> host visit <a href="http://localhost:8080">http://localhost:8080</a></p>
				<p>To test Cross-Origin requests form an <strong>Non-Allowed</strong> host visit <a href="http://127.0.0.1:8080/">http://127.0.0.1:8080/</a></p>
                <script>
                    document.getElementById('getButton').addEventListener('click', () => {
                        fetch('http://localhost:3000/hello')
                            .then(response => {
								if (response.headers.has('Access-Control-Allow-Origin')) {
									throw new Error('Access-Control-Allow-Origin header is present in a non-CORS request');
								}
								return response.text();
							})
                            .then(data => alert('Response from GET request: ' + data))
                            .catch(error => alert('Error:', error));
                    });
                    document.getElementById('postButton').addEventListener('click', () => {
                        fetch('http://localhost:3000/api/1', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({ key: 'value' }),
                        })
                            .then(response => {
                                if (response.headers.has('Access-Control-Allow-Origin')) {
									throw new Error('Access-Control-Allow-Origin header is present in a non-CORS request');
								}
								return response.json();
                            })
                            .then(data => alert('Response from POST request: ' + JSON.stringify(data)))
                            .catch(error => alert('Error:', error));
                    });
                    document.getElementById('optionsButton').addEventListener('click', () => {
                        fetch('http://localhost:3000/api/1', {
                            method: 'OPTIONS',
                        })
                            .then(response => {
                                if (!response.headers.has('Obvious-Header')) {
                                    throw new Error('Obvious-Header is missing, CORS middleware has intercepted the request');
                                }
								if (response.status !== 204) {
									throw new Error('Response status is not 204');
								}
								alert('Non-CORS OPTIONS request was successful');
							})
                            .catch(error => alert('Error:', error));
                    });
                </script>
            </body>
        </html>
    	`)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, API 👋!",
		})
	})

	app.Options("/*", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Obvious-Header", "Obvious-Value")
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Listen(":3000")
}

/*
To test the CORS policy, you can also use the following curl command:

curl -H "Origin: http://localhost:8080" \
     -H "Access-Control-Request-Method: GET" \
     -H "Access-Control-Request-Headers: Origin, Content-Type, Accept" \
     -X OPTIONS --verbose \
     http://localhost:3000
*/
