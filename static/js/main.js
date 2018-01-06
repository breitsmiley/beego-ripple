$(function () {

    var catHelper = {
        $mrCat: $(".mr-cat-wrapper"),
        $h1End: $("#h1-end"),
        $h1: $("h1"),
        catLeftMinLimit: 0,
        pageHeaderWidth: 0,
        catLeftMaxLimit: 0,
        catTop: 0,
        catLeft: 0,

        initSettings: function () {
            this.catLeftMinLimit = parseInt(this.$h1End.position().left) + 5;
            // this.pageHeaderWidth = this.$h1.width();
            this.catLeftMaxLimit = this.$h1.width() + this.$h1.position().left - 64;
            this.catTop = this.$h1.position().top;
            this.catLeft = this.getRandomArbitrary(this.catLeftMinLimit, this.catLeftMaxLimit);

            this.$mrCat.css({
                top: this.catTop +"px",
                left: this.catLeft+"px"
            });
        },

        getRandomArbitrary: function(min, max) {
            return Math.floor(Math.random() * (max - min) + min);
        },

        init: function () {

            setInterval(function () {
                if(catHelper.$mrCat.is(':hidden')) {
                    catHelper.initSettings();
                }
                catHelper.$mrCat.fadeToggle(5000,"linear");
            }, 7000);

        }
    };

    catHelper.init();

});

