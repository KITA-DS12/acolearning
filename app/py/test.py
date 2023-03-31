from parameters import Parameters
from graph import Graph

params = Parameters(num_ant=30, num_node=50, q=100, alpha=3, beta=5, rou=0.9,
                    max_iteration=300, min_tau=0, max_tau=3000,
                    border_unchanged=30, start_node=0)

graph = Graph(params, "")

print(graph.params.q)
