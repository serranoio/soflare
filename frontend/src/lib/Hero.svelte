<script lang="ts">
  import { CreateWebsiteLol } from "../util/model";
  import { TABS } from "../util/util"; 

  let website = "";
  
  const submit = async (e: any) => {
    e.preventDefault();

    const name = (document.querySelector("#website-name")! as HTMLInputElement).value 


    try {
        website = await CreateWebsiteLol(name);

    } catch(e) {

    } finally {

    }

  }

  </script>

<nav class="nav">
    <h3>Soflare</h3>
    <ul class="tabs-list">
        {#each TABS as tab}
        <li><a href={`#${tab}`}>{tab}</a></li>
        {/each}
    </ul>
    
</nav>

<main>
    <form>
        <h1>Create A Website!</h1>
        <div>

            <label for="name">Website Name</label>
            <input id="website-name" name="name" placeholder="website name"/>
        </div>
        <div>
            <button on:click={submit} type="submit">LETS PARTY BITCH</button>
        </div>
        <div>
            {#if website !== ""}
            <p>Visit website at! <a href={website}>{website}</a></p>
            {/if}
            <!-- {website === "" ? 
            nothing : -->
        <!-- } -->
        </div>
    </form>
</main>


<style>

main {
    margin: 0 auto;
    max-width: 130rem;
    height: calc(100vh - 80px);
}

button {
    font-size: 4rem;
    border-radius: 10px;
    padding: 1.2rem 2.4rem;
    letter-spacing: 2px;
    cursor: pointer;
    /* background-color: var(--yellow); */
    /* background-image: linear-gradient(var(--yellow), var(--green)); */
    /* color: var(--gray); */
    border: none;

    background: linear-gradient(
                90deg,
                var(--red),
                var(--orange),
                var(--yellow),
                var(--green),
                var(--red)
              ) 0 0 / 400% 100%;

  /* color: transparent; */
  /* background-clip: text; */
  /* -webkit-background-clip: text; */
  animation: move-bg 8s infinite linear;
}

button:hover {
}

@keyframes move-bg {
    to {
      background-position: 400% 0;
    }
}



form div {
    display: flex;
    justify-content: center;
    margin-bottom: 2.4rem;
    align-items: center;
    font-size: 4rem;
}

label {
    /* font-size: 2rem; */
    margin-right: 2rem;
    color: var(--orange);
}
input {
    font-size: 4rem;
    border-radius: 10px;
    /* height: 2rem; */
    padding: .6rem;
    border: 2px solid var(--orange);
    color: var(--orange);
    background-color: transparent;
}
h1 {
    font-size: 6rem;
    text-align: center;
    margin-bottom: calc(30vh);
}

nav {
    font-size: 3rem;
    height: 80px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 2.4rem;
}

h3 {
    /* font-size: 2.4rem; */
}

.tabs-list {
    display: flex;
    justify-content: space-between;
    gap: 2.4rem;
    list-style: none;
    color: var(--green);
}

a {
    text-decoration: none;
    color: var(--green);
    /* font-size: 2rem; */
}

</style>