const userNameFromURL = () => location.href.match("https://www.instagram.com/(.+)")[1].replace("/", "");

const showFollowers = username => {
    return document.querySelectorAll(`a[href='/${username}/followers/']`)[0].click()
};

const getFollowings = () => {
    return document.querySelectorAll("a.FPmhX._0imsa")
};

const algorithm = (index = 0, followingArchitects = {}) => {
    const username = userNameFromURL();

    if (!followingArchitects[username]) {
        followingArchitects[username] = []
    }

    setTimeout(() => {
        console.log(getFollowings())
        if (!getFollowings().length) {
            showFollowers(username);
        }

        setTimeout(() => {
            const followings = getFollowings() || [];
            const followingsCount = followings.length;

            // graphMap.push(username);

            followings[index].click();
            const biography = window._sharedData.entry_data.ProfilePage[0].graphql.user.biography;

            setTimeout(() => {
                window.history.back();
                algorithm(index + 1)
            }, 3000)

        }, 1000)
    }, 1000)
};