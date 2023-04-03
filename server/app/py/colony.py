from copy import deepcopy

from ant import Ant


class Colony:
    """
    Antの管理を行う

    Attributes
    ----------
    params: Parameters
        ACOを設定するすべてのパラメータを保存する
    graph: Graph
        ノードの座標やエッジの距離，フェロモン量などを保存する
    ants: list of Ant
        num_ant個のAntインスタンスを格納する
    """

    def __init__(self, parameters, graph):
        """
        Parameters
        ----------
        parameters: Parameters
            ACOを設定するすべてのパラメータを保存する
        graph: Graph
            ノードの座標やエッジの距離，フェロモン量などを保存する
        """
        self.params = parameters
        self.graph = graph
        self.ants = [Ant(self.params, self.graph)
                     for _ in range(self.params.num_ant)]

    def update_colony(self):
        """
        Antインスタンスのルートを構築し，
        通ったルートのフェロモン量をgraph.next_pheromone_edgeに適用する
        """
        self.__construct_ants()
        self.__calc_passed_pheromone()

    def get_best_ant_routes(self):
        """
        最も最適なアリのパスを取得する
        Returns
        -------
        [length, route]: [int, list of int]
            最適なアリのルートの長さとルート
        """
        length = 1e20
        route = []
        for ant in self.ants:
            if ant.calc_length_routes() < length:
                length = ant.calc_length_routes()
                route = deepcopy(ant.visited_routes)

        return [length, route]

    def reset_colony(self):
        for ant in self.ants:
            ant.reset_ant()

    def __construct_ants(self):
        """
        antsに格納したAntインスタンスのルートを構築する
        """
        for ant in self.ants:
            ant.construct_routes()

    def __calc_passed_pheromone(self):
        """
        Antインスタンスが通ったルートのフェロモン量を計算し、
        graph.next_pheromone_edgeを更新する
        """
        for ant in self.ants:
            ant.calc_passed_pheromone()
