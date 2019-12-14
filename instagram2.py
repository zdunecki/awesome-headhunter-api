import random
import argparse
from instapy import InstaPy
from instapy import smart_run
from instapy import util

from utils import is_architect, append_to_file, load_ig

insta_username = "lionel.maus"
insta_password = "qVL8fT3e39zSAF6"

session = InstaPy(
    username=insta_username,
    password=insta_password,
    headless_browser=False
)

visited_ig = []
not_valid_ig = []

load_ig("architects.csv", visited_ig)
load_ig("not_valid_architects.csv", not_valid_ig)


def grab_following(username):
    return session.grab_following(
        username=username,
        amount="full",
        live_match=True,
        store_locally=True
    )


def predict_architect_on_bio(username):
    session.validate_user_call(user_name=username)
    bio = util.getUserData("graphql.user.biography", session.browser)

    return is_architect(bio=bio)


def ig_link_id(username):
    return "https://instagram.com/" + username


def following_network(username):
    follow_is_valid = predict_architect_on_bio(username)

    if follow_is_valid:
        if username != args.root_user:
            followers_count = util.getUserData("graphql.user.edge_followed_by.count", session.browser)
            visited_ig.append(ig_link_id(username))
            append_to_file(
                file_name="architects.csv",
                data=[ig_link_id(username), followers_count],
                default_row=["ig_link", "followers"]
            )
    else:
        not_valid_ig.append(ig_link_id(username))
        append_to_file(
            file_name="not_valid_architects.csv",
            data=[ig_link_id(username)],
            default_row=["ig_link"]
        )
        return

    random_following = []
    following = grab_following(username)

    for follow in following:
        random_following.append(follow)

    random.shuffle(random_following)

    for follow in random_following:
        # TODO: not always because if we stop script and want to check following network of this username
        if ig_link_id(follow) in visited_ig:
            continue
        if ig_link_id(follow) in not_valid_ig:
            continue

        following_network(follow)


parser = argparse.ArgumentParser()

parser.add_argument(
    '-ru', '--root-user',
    required=True,
    help="root user name"
)

args = parser.parse_args()

with smart_run(session):
    following_network(args.root_user)
