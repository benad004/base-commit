package handler

import (
	"fmt"
	"net/http"
)

// Handler serves the main page with Vercel Analytics
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Base Commit - Vercel Analytics Demo</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 40px 20px;
            line-height: 1.6;
            color: #333;
        }
        h1 {
            color: #000;
            border-bottom: 2px solid #000;
            padding-bottom: 10px;
        }
        .info {
            background: #f5f5f5;
            padding: 20px;
            border-radius: 8px;
            margin: 20px 0;
        }
        code {
            background: #e0e0e0;
            padding: 2px 6px;
            border-radius: 3px;
            font-family: 'Courier New', monospace;
        }
    </style>
</head>
<body>
    <h1>Base Commit Project</h1>
    <p>This is a Go-based serverless function deployed on Vercel with Web Analytics enabled.</p>
    
    <div class="info">
        <h2>Vercel Web Analytics</h2>
        <p>This project is configured with <code>@vercel/analytics</code> to track page views and performance metrics.</p>
        <p>Analytics will be collected when deployed to Vercel and enabled in the dashboard.</p>
    </div>
    
    <h2>Features</h2>
    <ul>
        <li>Go serverless functions</li>
        <li>Vercel Web Analytics integration</li>
        <li>Automatic page view tracking</li>
        <li>Performance monitoring</li>
    </ul>

    <!-- Vercel Analytics Script -->
    <script>
        window.va = window.va || function () { (window.vaq = window.vaq || []).push(arguments); };
    </script>
    <script defer src="/_vercel/insights/script.js"></script>
</body>
</html>`

	fmt.Fprint(w, html)
}
