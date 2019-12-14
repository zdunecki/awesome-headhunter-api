import json
import re
import csv
import os
from pathlib import Path

from langdetect import detect

poland_stemming = ["Polska", "Polsce", "Polski", "Polish", "Poland"]
architect_stemming = [
    "Architekt",
    "Architekci",

    "Architekt Wnętrz",
    "Architekci Wnętrz",

    "Wnętrz",
    "Projektant Wnętrz",
    "Projektanci Wnętrz",

    "Architect",
    "Architects",
    "Interior Designer",
    "Interior Designers",
    "Interior Design"
]
most_popular_polish_cities = [
    "Warszawa",
    "Warsaw",
    "Cracow",
    "Krakow",
    "Kraków",
    "Wroclaw",
    "Wrocław",
    "Poznań",
    "Poznan",
    "Poznań",
    "Lodz",
    "Łódź",
    "Gdańsk",
    "Gdansk"
]
architect_keywords = architect_stemming + poland_stemming + most_popular_polish_cities

polish_website_regexp = "(?:(?:https):\/\/)?[\w/\-?=%.]+\.[pl]+"


def poland_probability(text):
    bio_lang = detect(text)

    if bio_lang == "pl":
        return True

    if bio_lang == "en":
        has_polish_website = re.search(polish_website_regexp, text)

        if has_polish_website:
            return True

    return False


def is_architect(bio):
    lower_format_bio = bio.lower()

    if not lower_format_bio:
        return False

    is_pl = poland_probability(lower_format_bio)

    if not is_pl:
        return False

    if not any(keyword.lower() in lower_format_bio for keyword in architect_keywords):
        return False

    return True


def load_ig(file_name, arr):
    with open(file_name) as f:
        cf = csv.DictReader(f, fieldnames=["ig_link"])

        for index, row in enumerate(cf):
            if index == 0:
                continue

            arr.append(row["ig_link"])


def append_to_file(file_name, data, default_row):
    exists = os.path.exists(file_name)

    with open(file_name, "a") as newFile:
        new_file_writer = csv.writer(newFile)
        if not exists:
            new_file_writer.writerow(default_row)

        new_file_writer.writerow(data)


#            ___A___
#           /       \
#          C         D
#        / | \     / | \
#       P  R  L   F  Q  S
#         / \       / \
#        O   E     G   H
#                 / \
#                N   M
#


def progressive_bfs(graph=None, visited=None, start="", fetch_neighbours=None):
    # initial fill graph
    if graph is None:
        graph = {}
        neighbours = fetch_neighbours(start)
        graph[start] = neighbours

    if visited is None:
        visited = []

    queue = [start]

    while queue:
        node = queue.pop(0)
        if node not in visited:
            visited.append(node)
            neighbours = graph[node]

            for neighbour in neighbours:
                if neighbour in visited:
                    continue

                queue.append(neighbour)
                new_neighbours = fetch_neighbours(neighbour)
                graph[neighbour] = new_neighbours

    return visited


def load_graph(root_user_name):
    home = str(Path.home())
    relationships_folder = os.path.join(home, "InstaPy/logs/" + root_user_name + "/relationship_data")
    user_names = os.listdir(relationships_folder)

    graph = {}

    for user_name in user_names:
        graph[user_name] = []
        user_followings_folder = os.path.join(relationships_folder, user_name, "following")

        for following_file in os.listdir(user_followings_folder):
            f = open(user_followings_folder + "/" + following_file, "r")
            data = json.load(f)
            # TODO: unique
            graph[user_name] = graph[user_name] + data

    if not graph:
        return None

    return graph
