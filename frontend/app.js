document.getElementById('shortenForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const longUrl = document.getElementById('longUrl').value;
    const customShort = document.getElementById('customShort').value;
    
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error');
    const linkAnchor = document.getElementById('shortenedLink');
    const metricsDiv = document.getElementById('metrics');

    // Reset visibility before making the new call
    resultDiv.classList.add('hidden');
    errorDiv.classList.add('hidden');

    try {
        const response = await fetch('http://localhost:3000/api/v1', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                url: longUrl,
                short: customShort
            })
        });

        const data = await response.json();

        if (!response.ok) {
            // Automatically catches 403 (already in use), 500, or 429 errors from Go
            throw new Error(data.error || 'Failed to shorten URL');
        }

        // Clean up the URL format to make sure it links correctly
        let rawShortUrl = data.short;
        if (!rawShortUrl.startsWith('http://') && !rawShortUrl.startsWith('https://')) {
            rawShortUrl = 'http://' + rawShortUrl;
        }

        // Update the UI link elements
        linkAnchor.href = rawShortUrl;
        linkAnchor.textContent = data.short;
        
        // Render your custom live Redis rate limit metrics!
        metricsDiv.innerHTML = `
            <p style="margin: 15px 0 0 0; font-size: 0.85rem; color: #a1a1aa;">
                ⚡ API Quota Remaining: <strong>${data.rate_limit}</strong> requests
            </p>
        `;
        
        resultDiv.classList.remove('hidden');

    } catch (err) {
        errorDiv.textContent = `${err.message}`;
        errorDiv.classList.remove('hidden');
    }
});