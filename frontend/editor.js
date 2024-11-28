import 'quill/dist/quill.core.css'
import 'quill/dist/quill.bubble.css'
import './style.css'

import Quill from 'quill';
import i18next from 'i18next';

var editor;

function save(event) {
    const titleEl = document.querySelector('#title')
    const authorEl = document.querySelector('#author')
    console.log(titleEl.value)
    console.log(authorEl.value)
    console.log(editor.getHtml())
    fetch("/api/save",
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: "POST",
            body: JSON.stringify(
                {
                    title: titleEl.value,
                    author: authorEl.value,
                    html: editor.getHtml()
                }
            )
        })
        .then(async function (res) {
            console.log(res)
            const receivedData = await res.json()
            console.log(receivedData)
            localStorage.setItem(receivedData["slug"], receivedData["secret"]);
            window.location.href="/"+receivedData["slug"]
        })
        .catch(function(res){ console.log(res) })
}

function preInit() {
    i18next.init({
        lng: 'en', // if you're using a language detector, do not define the lng option
        debug: true,
        resources: {
            en: {
                translation: {
                    "Title": "Title",
                    "Author": "author",
                    "TextPlaceholder": "Text...",
                    "Publish": "Publish",
                }
            },
            ru: {
                translation: {
                    "Title": "Заголовок",
                    "Author": "автор",
                    "TextPlaceholder": "Текст...",
                    "Publish": "Опубликовать",
                }
            },

        }
    });
}
function postInit() {
    Quill.prototype.getHtml = function () {
        return this.container.firstChild.innerHTML;
    };
    editor = setupQuill(document.querySelector('#editor'));
    document.querySelector('#publish_button').addEventListener("click", save);
}

function setupQuill(element) {
    const quill = new Quill(
        element,
        {
            debug: 'info',
            theme: 'bubble',
            placeholder: i18next.t("TextPlaceholder"),
        }
    );
    return quill
}

// APP
preInit()

document.querySelector('#app').innerHTML = `
<div class="article-header">
        <div class="article-header-title">
            <input type="text" id="title" placeholder="${i18next.t("Title")}" speech>
            <div class="input-title">${i18next.t("Title")}</div>
        </div>
        <div class="article-header-author">
            <input type="text" id="author" placeholder="${i18next.t("Author")}" speech6t>
            <div class="input-title">${i18next.t("Author")}</div>
        </div>
</div> 
<div id="editor"></div>
<div class="article-buttons">
   <button type="button" id="publish_button">${i18next.t("Publish")}</button>
</div>
`
postInit();

