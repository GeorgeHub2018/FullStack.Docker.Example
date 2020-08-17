sendRequest();

function sendRequest() {
    const Http = new XMLHttpRequest();
    const url='http://localhost:38080/World';
    Http.open("GET", url);
    Http.send();

    Http.onreadystatechange = (e) => {
        setResponse(Http.responseText);
    };
};

function setResponse(response) {
    var textNode = document.getElementById("response_text");
    if (textNode == null) {return};

    textNode.innerHTML = response;
};