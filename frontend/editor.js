import 'quill/dist/quill.core.css'
import 'quill/dist/quill.bubble.css'
import './style.css'

import Quill from 'quill';
import i18next from 'i18next';

function init() {
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

init()
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
   <button type="button">${i18next.t("Publish")}</button>
</div>
`

const editor = setupQuill(document.querySelector('#editor'));
