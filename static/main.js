const sendButton = document.getElementById("send-url")

const buttonChangeTheme = function() {
    let ready = false;
    return function(event){
        event.preventDefault();
        if(inputUrl.value.length > 0){
            if(!ready){
                ready = true;
                sendButton.className = "yellow-button";
            }
        } else {
            if(ready){
                ready = false;
                sendButton.className = "default-button";
            }
        }
    }
}();

const inputUrl = document.getElementById("input-url");
inputUrl.addEventListener("input", buttonChangeTheme);

const cloud = document.getElementById("cloud");
cloud.addEventListener("click", function(event){
    event.preventDefault();
    cloud.style.display = "none";
})

function showCloud(text) {
    cloud.style.display = "block";
    cloud.innerText = text;
}

const outputURL = document.getElementById("output-url");

document.getElementById("copy-button").addEventListener("click", function(event){
    event.preventDefault();
    navigator.clipboard.writeText(outputURL.innerText);
    showCloud("Успешно скопировано!");
})

const shareButton = document.getElementById("share-button");
shareButton.addEventListener("click", function(event){
    event.preventDefault();
    showCloud("Поделитесь ссылкой!");
})

const showWindow = function(){
    let ready = false;
    return function(){
        if(!ready){
            ready = true;
            document.getElementById("window-result").style.display = "inline-flex";
        }
    }
}();

const redirect = document.getElementById("redirect");

document.getElementById("send-url").addEventListener("click", function(event){
    event.preventDefault();
    console.log(inputUrl.value)
    fetch("./getShortUrl", {
        method: "POST",
        headers: {
            "content-type": "application/json",
        },
        body: JSON.stringify({"url": inputUrl.value}),
    }).
    then(response => response.json()).
    then(data => {
        showWindow();
        outputURL.innerText = data.url;
        redirect.href = data.url;
    }).catch(err => console.log(err));
});
