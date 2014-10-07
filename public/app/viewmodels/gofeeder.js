define(['plugins/http', 'durandal/app', 'knockout'], function(http, app, ko) {
    var feeder = {
        displayName: 'Gofeeder',
        feeds: ko.observableArray([]),
        items: ko.observableArray([]),
        activeFeed: ko.observable({}),
        activeItem: ko.observable({}),
        activate: function() {
            console.log('gofeeder activate called');
            if (this.items().length > 0) {
                return;
            }

            var self = this;
            return http.get('/feed/index').then(function(feeds) {
                console.log('response %o', feeds);
                self.feeds(feeds);
            });
        },
        setFeed: function(feed) {
            var self = this;
            this.activeFeed(feed);
            this.activeItem({});
            console.log('select called feed %o this %o ', feed, this);
            return http.get('/item/index?feedId=' + feed.Id).then(function(items) {
                self.items(items);
            });
        },
        setItem: function(item) {
            var self = this;
            this.activeItem(item);
            console.log('setting active item: %o', item);
            this.activeItem(item);
        }
    };

    /*feeder.getActiveItemName = ko.computed(function(){
        var name = (feeder.activeItem() && feeder.activeItem().Name) || '';
        return name;
    }, feeder);*/

    return feeder;
});
