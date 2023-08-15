document.getElementById('url-input').addEventListener('submit', submitForm);
document.getElementById('copy').addEventListener('click', copy);

let address = 'http://localhost:8080'
let text = document.getElementById('short-url-text');
let link = document.getElementById('short-url-link');
let copyButton = document.getElementById('copy');
copyButton.hidden = true;


function submitForm(event)
{
    event.preventDefault();
    const url = document.getElementById('url').value
    if (url.startsWith('http://') || url.startsWith('https://')) {
        text.innerText = "";
        fetch(address+'/getShortUrl', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({"url": url})
        }).then(response => response.json()).then(data => {
            let responseUrl = data.url;
            link.style.color = "black";
            text.innerText = responseUrl;
            link.href = responseUrl;
            copyButton.hidden = false;
        }).catch(_ => {
            invalidUrl();
        });
    } else {
        invalidUrl();
    }
}

function copy(){
    if (text.innerText.length > 0) {
        let err = navigator.clipboard.writeText(text.innerText);
        if (err) {
            console.log(err);
        }
    }
}

function invalidUrl(){
    link.style.color = "red";
    text.innerText = 'Invalid URL';
    link.href = '#';
    copyButton.hidden = true;
}