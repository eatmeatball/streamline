// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link RiParser}.
 */
public interface RiListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code Number}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterNumber(RiParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Number}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitNumber(RiParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(RiParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(RiParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(RiParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link RiParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(RiParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by {@link RiParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(RiParser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link RiParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(RiParser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link RiParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void enterBlockStatement(RiParser.BlockStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link RiParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void exitBlockStatement(RiParser.BlockStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BlockStms}
	 * labeled alternative in {@link RiParser#blockStatements}.
	 * @param ctx the parse tree
	 */
	void enterBlockStms(RiParser.BlockStmsContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BlockStms}
	 * labeled alternative in {@link RiParser#blockStatements}.
	 * @param ctx the parse tree
	 */
	void exitBlockStms(RiParser.BlockStmsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RiParser#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(RiParser.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link RiParser#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(RiParser.StartContext ctx);
}