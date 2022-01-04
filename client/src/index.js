// Test import of styles
import "@/styles/index.scss";
import { getSavedStocks, listenSocket, loadSymbols } from "@/js/stocks";

window.onload = () => {
  listenSocket();
  getSavedStocks();
};
