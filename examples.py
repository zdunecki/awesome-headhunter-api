from utils import is_architect, progressive_bfs

print(
    is_architect(
        "interior design, architecture HOLA DESIGN team is made of experienced designers and architects. In our portfolio we have hundreds of delivered interior and architecture projects. www.hola-design.pl")
)

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

data = {
    "A": ["C", "D"],
    "C": ["P", "R", "L"],
    "D": ["F", "Q", "S"],
    "R": ["O", "E"],
    "Q": ["G", "H"],
    "G": ["N", "M"],
}


def get_next(node):
    if node in data:
        return data[node]
    else:
        return []


print(
    progressive_bfs(graph=None, visited=None, start="A", fetch_neighbours=get_next)
)
