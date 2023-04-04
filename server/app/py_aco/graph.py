# import json
import math
import random


class Graph:
    """
    ノードやエッジ，フェロモンなどを管理する

    Attributes
    ----------
    params: Parameters
        ACOを設定するすべてのパラメータを保存する
    coordinates: list of tuple of float
        ノードの座標を格納する
    length_edge: list of int
        エッジの長さを格納する
    pheromone_edge: list of int
        現在蓄積されているエッジのフェロモン量を格納する
    next_pheromone_edge: list of int
        最新のサイクルで分泌されたフェロモン量を格納する
    heuristics_edge: list of int
        エッジのヒューリスティック値を格納する
    """

    def __init__(self, parameters, path):
        """
        Parameters
        ----------
        parameters: Parameters
            ACOを設定するすべてのパラメータを保存する
        """
        self.params = parameters
        num_node = self.params.num_node
        self.coordinates = [None] * num_node
        self.length_edge = [[0] * num_node for _ in range(num_node)]
        self.pheromone_edge = [[0] * num_node for _ in range(num_node)]
        self.next_pheromone_edge = [[0] * num_node for _ in range(num_node)]
        self.heuristics_edge = [[0] * num_node for _ in range(num_node)]
        self.__prepare_graph(path)

    def __prepare_graph(self, path):
        """
        受け取ったグラフ構造の情報を読み取る
        path: str
            グラフ構造を記述したファイルのパス
        """
        # with open(path, 'r') as f:
        #     file = json.load(f)
        # layouts = file.get("layouts", {})
        # nodes = layouts.get("nodes", {})
        # for node in nodes:
        #     node_id = node.get("id", 0)
        #     node_x = node.get("x", 0.0)
        #     node_y = node.get("y", 0.0)
        #     self.coordinates[node_id] = (node_x, node_y)
        # edges = layouts.get("edges", {})
        # for edge in edges:
        #     source = edge.get("source", 0)
        #     target = edge.get("target", 0)
        #     length = edge.get("length", 0)
        #     self.length_edge[source][target] = length
        #     self.length_edge[target][source] = length
        for i in range(self.params.num_node):
            self.coordinates[i] = (random.uniform(
                0, 1000), random.uniform(0, 1000))
        for i in range(self.params.num_node):
            for j in range(self.params.num_node):
                self.length_edge[i][j] = Graph.__calc_edge_length(
                    self.coordinates[i], self.coordinates[j])
        for i in range(self.params.num_node):
            for j in range(self.params.num_node):
                if i < j:
                    self.pheromone_edge[i][j] = self.params.q / \
                        self.length_edge[i][j]
                    self.pheromone_edge[j][i] = self.params.q / \
                        self.length_edge[i][j]
        for i in range(self.params.num_node):
            for j in range(self.params.num_node):
                if i < j:
                    h = self.params.q / self.length_edge[i][j]
                    self.heuristics_edge[i][j] = h
                    self.heuristics_edge[j][i] = h

    def reset_graph(self):
        self.next_pheromone_edge = [
            [0] * self.params.num_node for _ in range(self.params.num_node)]

    def reset_pheromones(self):
        """
        エッジに蓄積したフェロモンをすべてリセットする
        """
        for i in range(self.params.num_node):
            for j in range(self.params.num_node):
                if i < j:
                    self.pheromone_edge[i][j] = \
                        self.params.q / self.length_edge[i][j]
                    self.pheromone_edge[j][i] = \
                        self.params.q / self.length_edge[j][i]

    @staticmethod
    def __calc_edge_length(p1, p2):
        dx = (p1[0] - p2[0]) ** 2
        dy = (p1[1] - p2[1]) ** 2
        return math.sqrt(dx + dy)
