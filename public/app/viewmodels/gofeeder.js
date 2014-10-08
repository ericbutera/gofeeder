define(['plugins/http', 'durandal/app', 'durandal/system', 'knockout'], function(http, app, system, ko) {
    var feeder = {
        feeds: ko.observableArray([]),
        items: ko.observableArray([]),
        activeFeed: ko.observable({}),
        activeItem: ko.observable({}),
        activate: function() {
            /*
            this would prevent reloading feeds on subsequent navigations to the app
            if (this.feeds().length > 0) {
                return;
            }*/

            var self = this;
            return http.get('/feed/index').then(function(feeds) {
                system.log('fetched %o feeds', feeds.length);
                self.feeds(feeds);
            });
        },
        setFeed: function(feed) {
            var self = this;
            this.activeFeed(feed);
            this.activeItem({});
            return http.get('/item/index?feedId=' + feed.Id).then(function(items) {
                system.log('fetched %o items', items.length);
                self.items(items);
            });
        },
        setItem: function(item) {
            this.activeItem(item);
        },
        isFeedActive: function(feed) {
            return this.activeFeed() && this.activeFeed().Id == feed.Id;
        }
    };

    /*feeder.getActiveItemName = ko.computed(function(){
        var name = (feeder.activeItem() && feeder.activeItem().Name) || '';
        return name;
    }, feeder);*/

    return feeder;
});

