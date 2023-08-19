<script>
  import { Search } from '../wailsjs/go/main/App.js';
  import logo from './assets/images/tricolor_haze.jpg';

  let resultText = "Should the use of the word 'data' be considered heretical when it is used as a singular noun? (Y/N)"
  let query

  function search() {
    Search(query).then(result => {
      if (query.toUpperCase() === 'Y') {
        resultText = "Comrade, we shall patriotically correct your document with proper grammar :)";
      } else if (query.toUpperCase() === 'N') {
        resultText = "I wish I could as careless as you about using the English language.";
      } else {
        resultText = "Invalid input. Please enter 'Y' or 'N'.";
      }
    });
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={query} class="input" id="query" type="text"/>
    <button class="btn" on:click={search}>Search</button>
  </div>
</main>

<style>
    #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
