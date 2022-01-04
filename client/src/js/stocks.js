const stocks = [];
let searchedStocks = [];

const inputBox = document.getElementById("coin-search");
const searchWrapper = document.querySelector(".search-input");
const suggBox = searchWrapper.querySelector(".autocom-box");
let ws;
const stockMap = new Map();

export function listenSocket() {
    ws = new WebSocket(`${window.location.protocol.endsWith('s:') ? 'wss://' : 'ws://'}${window.location.host}/api/ws/ticker/`);
    ws.onmessage = (ev) => {
        const data = JSON.parse(ev.data);

        Object.keys(data).forEach((key) => {
            if (stockMap.has(key)) {
                const stock = stockMap.get(key);
                stock.value = data[key];
                stock.update();
            }
        });
        saveStocks();
    };

    ws.onopen = (ev) => {
        loadSymbols();
    };
}

class Stock {
    change = 0;
    name = "";
    value = 0;
    symbol_id = "";

    update() {
        const row = document.getElementById(this.symbol_id);
        if (row) {
            row.innerHTML = getStockString(this);
        }
        return `
        <tr id="${this.symbol_id}" class="stock ${
            this.change > 0 ? "increase" : "decrease"
        }">
            ${getStockString(this)}
        </tr>`;
    }
}

function getStockString(stock) {
    return `
            <td class="name"> ${stock.name} </td>
            <td class="value">${stock.value} </td>
            <td class="change"> ${(stock.change * stock.value).toFixed(2)} </td>
            <td class="percentage">${(stock.change * 100).toFixed(2)} %</td>
    `;
}

export function loadSymbols() {
    const savedStocks = getSavedStocks();

    savedStocks.forEach((row) => {
        const stock = new Stock();
        Object.assign(stock, row);
        const str = stock.update();
        $("table tbody").append(str);
        stocks.push(stock);
        stockMap.set(stock.symbol_id, stock);
    });

    ws.send(stocks.map((s) => s.symbol_id));
}

export function getSavedStocks() {
    const stocksSaved = localStorage.getItem("stocks");

    if (stocksSaved) {
        return JSON.parse(stocksSaved);
    }
    return [];
}

function saveStocks() {
    localStorage.setItem("stocks", JSON.stringify(stocks));
}

inputBox.onkeyup = (e) => {
    let userData = e.target.value; //user enetered data
    if (userData) {
        searchWrapper.classList.add("active"); //show autocomplete box
        getAllSymbols(userData);
    } else {
        searchWrapper.classList.remove("active"); //hide autocomplete box
    }
};

function showSuggestions(list) {
    const listData = document.createElement("ul");
    list.forEach((l, idx) => {
        const li = document.createElement("li");
        li.innerText = l;
        li.onclick = () => {
            selectSymbol(idx);
        };
        li.id = idx;
        listData.append(li);
    });
    suggBox.innerHTML = "";
    suggBox.append(listData);
}

function getAllSymbols(query) {
    let url = "/api/static_data/search";
    if (query) {
        url += "?query=" + query;
    }
    fetch(url)
        .then((data) => data.json())
        .then((data) => {
            searchedStocks = data;
            showSuggestions(searchedStocks.map((s) => s.Name));
        });
}

function selectSymbol(idx) {
    const stock = new Stock();
    stock.name = searchedStocks[idx].Name;
    stock.value = searchedStocks[idx]["LTP"];
    stock.symbol_id = searchedStocks[idx]["SymbolID"];
    const str = stock.update();
    $("table tbody").append(str);
    searchWrapper.classList.remove("active");
    stocks.push(stock);
    saveStocks();
    stockMap.set(stock.symbol_id, stock);

    ws.send(searchedStocks[idx].SymbolID);
}
