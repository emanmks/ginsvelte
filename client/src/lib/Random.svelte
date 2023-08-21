<script lang="ts">
    let random: number = -1;

    function getRandom() {
        fetch('./api/random')
            .then(res => res.json())
            .then(text => random = parseFloat(text.random));
    }
</script>

<h2>Ypur random digit is {random}</h2>
<button on:click={getRandom}>
    Get a new random digit
</button>
