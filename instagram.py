import argparse
from instapy import InstaPy
from instapy import smart_run

from utils import progressive_bfs, load_graph

insta_username = "lionel.maus"
insta_password = "qVL8fT3e39zSAF6"

session = InstaPy(
    username=insta_username,
    password=insta_password,
    headless_browser=False
)

parser = argparse.ArgumentParser()

parser.add_argument(
    '-ru', '--root-user',
    required=True,
    help="root user name"
)

args = parser.parse_args()


def get_followings(username):
    return session.grab_following(
        username=username,
        amount="full",
        live_match=True,
        store_locally=True
    )


with smart_run(session):
    graph = load_graph(insta_username)

    visited = None

    if graph:
        visited = []
        for user_name in graph:
            if user_name != args.root_user:
                visited.append(user_name)

    progressive_bfs(graph=graph, visited=visited, start=args.root_user, fetch_neighbours=get_followings)
