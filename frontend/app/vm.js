var am = new Vue({
    el: '#app',
    data: {
        individual: true,
        notary: true,
        agency: false,

        loading: false,

        postcode: "",
        saleType: "",

        results: []
    },
    methods: {
        lookup() {
            this.loading = true
            axios.get(`/api/lookup?postcode=${this.postcode}&saleType=${this.saleType}&individual=${this.individual ? "true" : ""}&notary=${this.notary ? "true" : ""}&agency=${this.agency ? "true" : ""}`)
            .then((res) => {
                this.results = res.data
                this.loading = false
            })
        },
        openBrowser(url) {
            console.log(url)
            window.goapp.openURL(url)
        }
    }
})