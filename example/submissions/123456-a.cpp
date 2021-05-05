#include <bits/stdc++.h>
using namespace std;


int searchVerticalTop(char grid[100][100], string kata, int n, int m,int posx, int posy) {

	if (posy < n) {
		string tmp = "";
		for(int i=posy; i<n; i++) {
			tmp += grid[i][posx];
			if(tmp == kata) {
				// cout << "searchVerticalTop" << '\n';
				return 1;
			}
		}
	}
	return 0;
}

int searchVerticalBot(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	if(posy > 0) {
		string tmp = "";
		for(int i=posy; i>=0; i--) {
			tmp += grid[i][posx];

			if(tmp == kata) {
				// cout << posx << " " << posy << '\n';
				return 1;
			}
		}

	}
	return 0;
}

int searchHorizontalRight(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	
	if (posx < m) {
		string tmp = "";
		for(int i=posx; i<m; i++) {
			tmp += grid[posy][i];
			if(tmp == kata) {
				// cout << "searchHorizontalRight" << '\n';
				return 1;
			}
		}
	}
	return 0;
}

int searchHorizontalLeft(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	
	if (posx > 0) {
		string tmp = "";
		for(int i=posx; i>=0; i--) {
			tmp += grid[posy][i];
			if(tmp == kata) {
				// cout << posx << " " << posy << '\n';
				return 1;
			}
		}
	}
	return 0;
}

int searchDiagonalTopRight(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	
	string tmp = "";
	while(posx < m && posy >= 0) {
		tmp += grid[posy][posx];
		posy--;
		posx++;
		if(tmp == kata) {
			// cout << posx << " " << posy << '\n';

			return 1;
		}
	}
	return 0;
}

int searchDiagonalTopLeft(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	string tmp = "";
	while(posx >= 0 && posy >= 0) {
		tmp += grid[posy][posx];
		posy--;
		posx--;
		if(tmp == kata) {
			// cout << posx << " " << posy << '\n';
			return 1;
		}
	}
	return 0;
}

int searchDiagonalBotLeft(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	
	string tmp = "";
	while(posx >= 0 && posy < n) {
		tmp += grid[posy][posx];
		posy++;
		posx--;
		if(tmp == kata) {
			// cout << posx << " " << posy << '\n';
			return 1;
		}
	}
	return 0;
}

int searchDiagonalBotRight(char grid[100][100], string kata, int n, int m,int posx, int posy) {
	string tmp = "";
	while(posx < m && posy < n) {
		tmp += grid[posy][posx];
		posy++;
		posx++;
		if(tmp == kata) {
			// cout << "searchDiagonalBotRight" << '\n';
			return 1;
		}
	}
	return 0;
}


int main() {
	int t, tcase = 1;
	int n, m;
	string w;

	cin >> t;
	while(t--) {
		cin >> n >> m;
		char grid[100][100];

		char tmp;
		for(int i=0; i<n; i++) {
			for(int j=0; j<m; j++) {
				cin >> tmp;
				grid[i][j] = tmp;
			}
		}

		cin >> w;
		// cout << searchVerticalBot(grid, w, n, m, 3, 2) << '\n';
		int result = 0;
		for(int i=0; i<n; i++) {
			for(int j=0; j<m; j++) {
				result += searchVerticalTop(grid, w, n, m, j, i);
				result += searchVerticalBot(grid, w, n, m, j, i);
				result += searchHorizontalRight(grid, w, n, m, j, i);
				result += searchHorizontalLeft(grid, w, n, m, j, i);
				result += searchDiagonalTopRight(grid, w, n, m, j, i);
				result += searchDiagonalTopLeft(grid, w, n, m, j, i);
				result += searchDiagonalBotRight(grid, w, n, m, j, i);
				result += searchDiagonalBotLeft(grid, w, n, m, j, i);
			}
		}

		cout << "Case " <<tcase++ << ": " << result << '\n';

	}

}