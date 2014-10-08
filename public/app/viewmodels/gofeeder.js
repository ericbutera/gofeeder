define(['plugins/http', 'durandal/app', 'durandal/system', 'knockout'], function(http, app, system, ko) {
    // use valueHasMutated when modifying observableArrays manually and need to update subscribers
    // make some sort of grid for items that uses a partial item list
    // make grid expand into full detail by fetching /item/view?id=
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
                if (feeds && feeds.length) {
                    system.log('fetched %o feeds', feeds.length);
                    self.feeds(feeds);
                }
            });
        },
        setFeed: function(feed) {
            var self = this;
            this.activeFeed(feed);
            this.activeItem({});
            return http.get('/item/index?feedId=' + feed.Id).then(function(items) {
                if (items && items.length) {
                    system.log('fetched %o items', items.length);
                    self.items(items);
                } else {
                    self.items({});
                }
            });
        },
        setItem: function(listItem) {
            var self = this;
            system.log('setting list item %o', listItem);
            return http.get('/item/view?id=' + listItem.Id).then(function(item) {
                if (item && item.Id) {
                    system.log('setting item %o', item);
                    self.activeItem(item);
                }
            });
        },
        isFeedActive: function(feed) {
            return this.activeFeed() && this.activeFeed().Id == feed.Id;
        }
    };

    /*feeder.getActiveItemName = ko.computed(function(){
        var name = (feeder.activeItem() && feeder.activeItem().Name) || '';
        return name;
    }, feeder);*/

    system.log("feeder %o", feeder);
    return feeder;
});

